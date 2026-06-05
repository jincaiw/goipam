package api

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/domain/scan"
	"github.com/jincaiw/goipam/internal/repository"
)

// AddressHandler handles address API endpoints.
type AddressHandler struct {
	addressRepo *repository.AddressRepository
	subnetRepo  *repository.SubnetRepository
}

// NewAddressHandler creates a new AddressHandler.
func NewAddressHandler(addressRepo *repository.AddressRepository, subnetRepo *repository.SubnetRepository) *AddressHandler {
	return &AddressHandler{
		addressRepo: addressRepo,
		subnetRepo:  subnetRepo,
	}
}

// List handles GET /api/v1/addresses
func (h *AddressHandler) List(w http.ResponseWriter, r *http.Request) {
	subnetIDStr := r.URL.Query().Get("subnet_id")
	if subnetIDStr == "" {
		BadRequest(w, r, errcode.ValidationError, "subnet_id query parameter is required")
		return
	}

	subnetID, err := strconv.ParseInt(subnetIDStr, 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet_id")
		return
	}

	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	addresses, err := h.addressRepo.List(r.Context(), subnetID, int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list addresses")
		return
	}

	total, err := h.addressRepo.Count(r.Context(), subnetID)
	if err != nil {
		InternalError(w, r, "failed to count addresses")
		return
	}

	OKPaged(w, r, addresses, total, page, pageSize)
}

// Get handles GET /api/v1/addresses/{id}
func (h *AddressHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid address id")
		return
	}

	address, err := h.addressRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "address")
		return
	}

	OK(w, r, address)
}

// Create handles POST /api/v1/addresses
func (h *AddressHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SubnetID    int64  `json:"subnet_id"`
		IP          string `json:"ip"`
		Hostname    string `json:"hostname"`
		Mac         string `json:"mac"`
		Description string `json:"description"`
		Status      string `json:"status"`
		TagID       *int64 `json:"tag_id"`
		IsGateway   bool   `json:"is_gateway"`
		DeviceID    *int64 `json:"device_id"`
		DevicePort  string `json:"device_port"`
		Owner       string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.IP == "" {
		BadRequest(w, r, errcode.ValidationError, "ip is required")
		return
	}

	if req.SubnetID == 0 {
		BadRequest(w, r, errcode.ValidationError, "subnet_id is required")
		return
	}

	// Get the subnet to validate IP is within range
	subnet, err := h.subnetRepo.GetByID(r.Context(), req.SubnetID)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	// Validate that IP is within the subnet's CIDR range
	isInSubnet, err := ipam.IsSubnetOf(subnet.Cidr, req.IP+"/32")
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid IP or CIDR")
		return
	}
	if !isInSubnet {
		Error(w, r, http.StatusBadRequest, errcode.AddressNotInSubnet, "IP address is not within the subnet range")
		return
	}

	// Convert IP to binary for storage
	ipBytes, err := ipam.IPToBytes(req.IP)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid IP address")
		return
	}

	// Check if IP is already occupied
	existing, err := h.addressRepo.GetByIPText(r.Context(), req.SubnetID, req.IP)
	if err == nil && existing != nil {
		Error(w, r, http.StatusConflict, errcode.AddressOccupied, "IP address is already in use")
		return
	}

	if req.Status == "" {
		req.Status = "active"
	}

	id, err := h.addressRepo.Create(
		r.Context(), req.SubnetID, ipBytes, req.IP,
		req.Hostname, req.Mac, req.Description, req.Status,
		req.TagID, req.IsGateway, req.DeviceID, req.DevicePort, req.Owner,
	)
	if err != nil {
		InternalError(w, r, "failed to create address")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// FirstFree handles POST /api/v1/addresses/first-free
func (h *AddressHandler) FirstFree(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SubnetID    int64  `json:"subnet_id"`
		Hostname    string `json:"hostname"`
		Mac         string `json:"mac"`
		Description string `json:"description"`
		Status      string `json:"status"`
		TagID       *int64 `json:"tag_id"`
		IsGateway   bool   `json:"is_gateway"`
		DeviceID    *int64 `json:"device_id"`
		DevicePort  string `json:"device_port"`
		Owner       string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.SubnetID == 0 {
		BadRequest(w, r, errcode.ValidationError, "subnet_id is required")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), req.SubnetID)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	existingIPs, err := h.addressRepo.ListIPsBySubnet(r.Context(), req.SubnetID)
	if err != nil {
		InternalError(w, r, "failed to list existing IPs")
		return
	}

	ip, err := ipam.GetFirstAvailableIP(subnet.Cidr, existingIPs)
	if err != nil {
		Error(w, r, http.StatusNotFound, errcode.NoAvailableIP, err.Error())
		return
	}

	ipBytes, err := ipam.IPToBytes(ip)
	if err != nil {
		InternalError(w, r, "failed to convert IP to bytes")
		return
	}

	if req.Status == "" {
		req.Status = "active"
	}

	id, err := h.addressRepo.Create(
		r.Context(), req.SubnetID, ipBytes, ip,
		req.Hostname, req.Mac, req.Description, req.Status,
		req.TagID, req.IsGateway, req.DeviceID, req.DevicePort, req.Owner,
	)
	if err != nil {
		InternalError(w, r, "failed to create address")
		return
	}

	Created(w, r, map[string]interface{}{
		"id": id,
		"ip": ip,
	})
}

// Update handles PUT /api/v1/addresses/{id}
func (h *AddressHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid address id")
		return
	}

	var req struct {
		Hostname    string `json:"hostname"`
		Mac         string `json:"mac"`
		Description string `json:"description"`
		Status      string `json:"status"`
		TagID       *int64 `json:"tag_id"`
		IsGateway   bool   `json:"is_gateway"`
		DeviceID    *int64 `json:"device_id"`
		DevicePort  string `json:"device_port"`
		Owner       string `json:"owner"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if err := h.addressRepo.Update(r.Context(), id, req.Hostname, req.Mac, req.Description, req.Status, req.TagID, req.IsGateway, req.DeviceID, req.DevicePort, req.Owner); err != nil {
		NotFound(w, r, "address")
		return
	}

	OK(w, r, map[string]string{"message": "address updated"})
}

// Delete handles DELETE /api/v1/addresses/{id}
func (h *AddressHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid address id")
		return
	}

	if err := h.addressRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "address")
		return
	}

	OK(w, r, map[string]string{"message": "address deleted"})
}

// Ping handles POST /api/v1/addresses/{id}/ping
func (h *AddressHandler) Ping(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid address id")
		return
	}

	address, err := h.addressRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "address")
		return
	}

	result := scan.PingProbe(r.Context(), address.IpText, 3*time.Second)

	OK(w, r, map[string]interface{}{
		"ip":      result.IP,
		"alive":   result.Alive,
		"message": result.Message,
	})
}

// DnsResolve handles POST /api/v1/addresses/{id}/dns-resolve
func (h *AddressHandler) DnsResolve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid address id")
		return
	}

	address, err := h.addressRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "address")
		return
	}

	names, err := net.LookupAddr(address.IpText)
	resolved := err == nil

	OK(w, r, map[string]interface{}{
		"ip":       address.IpText,
		"resolved": resolved,
		"names":    names,
	})
}

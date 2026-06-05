package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/rbac"
	"github.com/jincaiw/goipam/internal/repository"
)

// RequestHandler handles IP request API endpoints.
type RequestHandler struct {
	requestRepo *repository.RequestRepository
	subnetRepo  *repository.SubnetRepository
	addressRepo *repository.AddressRepository
}

// NewRequestHandler creates a new RequestHandler.
func NewRequestHandler(requestRepo *repository.RequestRepository, subnetRepo *repository.SubnetRepository, addressRepo *repository.AddressRepository) *RequestHandler {
	return &RequestHandler{
		requestRepo: requestRepo,
		subnetRepo:  subnetRepo,
		addressRepo: addressRepo,
	}
}

// List handles GET /api/v1/requests
func (h *RequestHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	requests, err := h.requestRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list IP requests")
		return
	}

	total, err := h.requestRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count IP requests")
		return
	}

	OKPaged(w, r, requests, total, page, pageSize)
}

// Create handles POST /api/v1/requests
func (h *RequestHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SubnetID    int64  `json:"subnet_id"`
		RequestedIP string `json:"requested_ip"`
		Hostname    string `json:"hostname"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.SubnetID == 0 {
		BadRequest(w, r, errcode.ValidationError, "subnet_id is required")
		return
	}

	// Verify subnet exists and allows requests
	subnet, err := h.subnetRepo.GetByID(r.Context(), req.SubnetID)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	if !subnet.AllowRequest {
		Error(w, r, http.StatusBadRequest, errcode.RequestNotAllowed, "this subnet does not allow IP requests")
		return
	}

	// If a specific IP was requested, check if it's already occupied
	if req.RequestedIP != "" {
		existing, err := h.addressRepo.GetByIPText(r.Context(), req.SubnetID, req.RequestedIP)
		if err == nil && existing != nil {
			Error(w, r, http.StatusConflict, errcode.RequestIPOccupied, "requested IP is already in use")
			return
		}
	}

	// Get applicant ID from auth context
	applicantID := rbac.GetUserID(r.Context())

	id, err := h.requestRepo.Create(
		r.Context(), req.SubnetID, applicantID,
		req.RequestedIP, req.Hostname, req.Description,
	)
	if err != nil {
		InternalError(w, r, "failed to create IP request")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Approve handles POST /api/v1/requests/{id}/approve
func (h *RequestHandler) Approve(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request id")
		return
	}

	var req struct {
		Comment string `json:"comment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	ipReq, err := h.requestRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "IP request")
		return
	}

	if ipReq.Status != "pending" {
		BadRequest(w, r, errcode.ValidationError, "request is not in pending status")
		return
	}

	approverID := rbac.GetUserID(r.Context())

	// Determine the IP to assign
	var assignIP string
	if ipReq.RequestedIP != "" {
		// Check if the requested IP is still available
		existing, err := h.addressRepo.GetByIPText(r.Context(), ipReq.SubnetID, ipReq.RequestedIP)
		if err == nil && existing != nil {
			Error(w, r, http.StatusConflict, errcode.RequestIPOccupied, "requested IP is already in use")
			return
		}
		assignIP = ipReq.RequestedIP
	} else {
		// Auto-assign the first available IP
		subnet, err := h.subnetRepo.GetByID(r.Context(), ipReq.SubnetID)
		if err != nil {
			InternalError(w, r, "subnet not found")
			return
		}

		existingIPs, err := h.addressRepo.ListIPsBySubnet(r.Context(), ipReq.SubnetID)
		if err != nil {
			InternalError(w, r, "failed to list existing IPs")
			return
		}

		ip, err := ipam.GetFirstAvailableIP(subnet.Cidr, existingIPs)
		if err != nil {
			Error(w, r, http.StatusNotFound, errcode.NoAvailableIP, err.Error())
			return
		}
		assignIP = ip
	}

	// Create the address
	ipBytes, err := ipam.IPToBytes(assignIP)
	if err != nil {
		InternalError(w, r, "invalid IP address")
		return
	}

	_, err = h.addressRepo.Create(
		r.Context(), ipReq.SubnetID, ipBytes, assignIP,
		ipReq.Hostname, "", ipReq.Description, "active",
		nil, false, nil, "", "",
	)
	if err != nil {
		InternalError(w, r, "failed to create address")
		return
	}

	// Update request status
	if err := h.requestRepo.UpdateStatus(r.Context(), id, "approved", &approverID, req.Comment); err != nil {
		InternalError(w, r, "failed to update request status")
		return
	}

	OK(w, r, map[string]interface{}{
		"message": "request approved",
		"ip":      assignIP,
	})
}

// Reject handles POST /api/v1/requests/{id}/reject
func (h *RequestHandler) Reject(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request id")
		return
	}

	var req struct {
		Comment string `json:"comment"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	ipReq, err := h.requestRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "IP request")
		return
	}

	if ipReq.Status != "pending" {
		BadRequest(w, r, errcode.ValidationError, "request is not in pending status")
		return
	}

	approverID := rbac.GetUserID(r.Context())

	if err := h.requestRepo.UpdateStatus(r.Context(), id, "rejected", &approverID, req.Comment); err != nil {
		InternalError(w, r, "failed to update request status")
		return
	}

	OK(w, r, map[string]string{"message": "request rejected"})
}

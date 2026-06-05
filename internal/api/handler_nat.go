package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// NatHandler handles NAT API endpoints.
type NatHandler struct {
	natRepo *repository.NatRepository
}

// NewNatHandler creates a new NatHandler.
func NewNatHandler(natRepo *repository.NatRepository) *NatHandler {
	return &NatHandler{natRepo: natRepo}
}

// List handles GET /api/v1/nats
func (h *NatHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	nats, err := h.natRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list nats")
		return
	}

	total, err := h.natRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count nats")
		return
	}

	OKPaged(w, r, nats, total, page, pageSize)
}

// Get handles GET /api/v1/nats/{id}
func (h *NatHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nat id")
		return
	}

	nat, err := h.natRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "nat")
		return
	}

	OK(w, r, nat)
}

// Create handles POST /api/v1/nats
func (h *NatHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name             string `json:"name"`
		Type             string `json:"type"`
		InsideAddressID  *int64 `json:"inside_address_id"`
		OutsideAddressID *int64 `json:"outside_address_id"`
		InsideIP         string `json:"inside_ip"`
		OutsideIP        string `json:"outside_ip"`
		InsidePort       *int32 `json:"inside_port"`
		OutsidePort      *int32 `json:"outside_port"`
		Protocol         string `json:"protocol"`
		DeviceID         *int64 `json:"device_id"`
		Description      string `json:"description"`
		Enabled          bool   `json:"enabled"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}
	if req.Type == "" {
		req.Type = "snat"
	}
	if req.Protocol == "" {
		req.Protocol = "tcp"
	}

	id, err := h.natRepo.Create(r.Context(), req.Name, req.Type, req.InsideAddressID, req.OutsideAddressID, req.InsideIP, req.OutsideIP, req.InsidePort, req.OutsidePort, req.Protocol, req.DeviceID, req.Description, req.Enabled)
	if err != nil {
		InternalError(w, r, "failed to create nat")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/nats/{id}
func (h *NatHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nat id")
		return
	}

	var req struct {
		Name             string `json:"name"`
		Type             string `json:"type"`
		InsideAddressID  *int64 `json:"inside_address_id"`
		OutsideAddressID *int64 `json:"outside_address_id"`
		InsideIP         string `json:"inside_ip"`
		OutsideIP        string `json:"outside_ip"`
		InsidePort       *int32 `json:"inside_port"`
		OutsidePort      *int32 `json:"outside_port"`
		Protocol         string `json:"protocol"`
		DeviceID         *int64 `json:"device_id"`
		Description      string `json:"description"`
		Enabled          bool   `json:"enabled"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	if err := h.natRepo.Update(r.Context(), id, req.Name, req.Type, req.InsideAddressID, req.OutsideAddressID, req.InsideIP, req.OutsideIP, req.InsidePort, req.OutsidePort, req.Protocol, req.DeviceID, req.Description, req.Enabled); err != nil {
		NotFound(w, r, "nat")
		return
	}

	OK(w, r, map[string]string{"message": "nat updated"})
}

// Delete handles DELETE /api/v1/nats/{id}
func (h *NatHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nat id")
		return
	}

	if err := h.natRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "nat")
		return
	}

	OK(w, r, map[string]string{"message": "nat deleted"})
}

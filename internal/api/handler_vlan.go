package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// VlanHandler handles VLAN and L2 domain API endpoints.
type VlanHandler struct {
	vlanRepo *repository.VlanRepository
}

// NewVlanHandler creates a new VlanHandler.
func NewVlanHandler(vlanRepo *repository.VlanRepository) *VlanHandler {
	return &VlanHandler{vlanRepo: vlanRepo}
}

// ListVlans handles GET /api/v1/vlans
func (h *VlanHandler) ListVlans(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	vlans, err := h.vlanRepo.ListVlans(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list vlans")
		return
	}

	total, err := h.vlanRepo.CountVlans(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count vlans")
		return
	}

	OKPaged(w, r, vlans, total, page, pageSize)
}

// GetVlan handles GET /api/v1/vlans/{id}
func (h *VlanHandler) GetVlan(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vlan id")
		return
	}

	vlan, err := h.vlanRepo.GetVlanByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "vlan")
		return
	}

	OK(w, r, vlan)
}

// CreateVlan handles POST /api/v1/vlans
func (h *VlanHandler) CreateVlan(w http.ResponseWriter, r *http.Request) {
	var req struct {
		L2DomainID  *int64 `json:"l2_domain_id"`
		VlanID      int32  `json:"vlan_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}
	if req.VlanID < 1 || req.VlanID > 4094 {
		BadRequest(w, r, errcode.ValidationError, "vlan_id must be between 1 and 4094")
		return
	}

	id, err := h.vlanRepo.CreateVlan(r.Context(), req.L2DomainID, req.VlanID, req.Name, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create vlan")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// UpdateVlan handles PUT /api/v1/vlans/{id}
func (h *VlanHandler) UpdateVlan(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vlan id")
		return
	}

	var req struct {
		L2DomainID  *int64 `json:"l2_domain_id"`
		VlanID      int32  `json:"vlan_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}
	if req.VlanID < 1 || req.VlanID > 4094 {
		BadRequest(w, r, errcode.ValidationError, "vlan_id must be between 1 and 4094")
		return
	}

	if err := h.vlanRepo.UpdateVlan(r.Context(), id, req.L2DomainID, req.VlanID, req.Name, req.Description); err != nil {
		NotFound(w, r, "vlan")
		return
	}

	OK(w, r, map[string]string{"message": "vlan updated"})
}

// DeleteVlan handles DELETE /api/v1/vlans/{id}
func (h *VlanHandler) DeleteVlan(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vlan id")
		return
	}

	if err := h.vlanRepo.DeleteVlan(r.Context(), id); err != nil {
		NotFound(w, r, "vlan")
		return
	}

	OK(w, r, map[string]string{"message": "vlan deleted"})
}

// ListL2Domains handles GET /api/v1/l2-domains
func (h *VlanHandler) ListL2Domains(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	domains, err := h.vlanRepo.ListL2Domains(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list l2 domains")
		return
	}

	total, err := h.vlanRepo.CountL2Domains(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count l2 domains")
		return
	}

	OKPaged(w, r, domains, total, page, pageSize)
}

// CreateL2Domain handles POST /api/v1/l2-domains
func (h *VlanHandler) CreateL2Domain(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	id, err := h.vlanRepo.CreateL2Domain(r.Context(), req.Name, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create l2 domain")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

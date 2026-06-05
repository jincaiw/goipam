package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// VrfHandler handles VRF API endpoints.
type VrfHandler struct {
	vrfRepo *repository.VrfRepository
}

// NewVrfHandler creates a new VrfHandler.
func NewVrfHandler(vrfRepo *repository.VrfRepository) *VrfHandler {
	return &VrfHandler{vrfRepo: vrfRepo}
}

// List handles GET /api/v1/vrfs
func (h *VrfHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	vrfs, err := h.vrfRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list vrfs")
		return
	}

	total, err := h.vrfRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count vrfs")
		return
	}

	OKPaged(w, r, vrfs, total, page, pageSize)
}

// Get handles GET /api/v1/vrfs/{id}
func (h *VrfHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vrf id")
		return
	}

	vrf, err := h.vrfRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "vrf")
		return
	}

	OK(w, r, vrf)
}

// Create handles POST /api/v1/vrfs
func (h *VrfHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Rd          string `json:"rd"`
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

	id, err := h.vrfRepo.Create(r.Context(), req.Name, req.Rd, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create vrf")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/vrfs/{id}
func (h *VrfHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vrf id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Rd          string `json:"rd"`
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

	if err := h.vrfRepo.Update(r.Context(), id, req.Name, req.Rd, req.Description); err != nil {
		NotFound(w, r, "vrf")
		return
	}

	OK(w, r, map[string]string{"message": "vrf updated"})
}

// Delete handles DELETE /api/v1/vrfs/{id}
func (h *VrfHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid vrf id")
		return
	}

	if err := h.vrfRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "vrf")
		return
	}

	OK(w, r, map[string]string{"message": "vrf deleted"})
}

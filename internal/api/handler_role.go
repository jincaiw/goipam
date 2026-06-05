package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// RoleHandler handles role API endpoints.
type RoleHandler struct {
	roleRepo *repository.RoleRepository
}

// NewRoleHandler creates a new RoleHandler.
func NewRoleHandler(roleRepo *repository.RoleRepository) *RoleHandler {
	return &RoleHandler{roleRepo: roleRepo}
}

// List handles GET /api/v1/roles
func (h *RoleHandler) List(w http.ResponseWriter, r *http.Request) {
	roles, err := h.roleRepo.List(r.Context())
	if err != nil {
		InternalError(w, r, "failed to list roles")
		return
	}

	OK(w, r, roles)
}

// Get handles GET /api/v1/roles/{id}
func (h *RoleHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid role id")
		return
	}

	role, err := h.roleRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "role")
		return
	}

	OK(w, r, role)
}

// Create handles POST /api/v1/roles
func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	id, err := h.roleRepo.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create role")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/roles/{id}
func (h *RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid role id")
		return
	}

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

	if err := h.roleRepo.Update(r.Context(), id, req.Name, req.Description); err != nil {
		NotFound(w, r, "role")
		return
	}

	OK(w, r, map[string]string{"message": "role updated"})
}

// Delete handles DELETE /api/v1/roles/{id}
func (h *RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid role id")
		return
	}

	if err := h.roleRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "role")
		return
	}

	OK(w, r, map[string]string{"message": "role deleted"})
}

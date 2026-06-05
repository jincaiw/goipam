package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// GroupHandler handles group API endpoints.
type GroupHandler struct {
	groupRepo *repository.GroupRepository
}

// NewGroupHandler creates a new GroupHandler.
func NewGroupHandler(groupRepo *repository.GroupRepository) *GroupHandler {
	return &GroupHandler{groupRepo: groupRepo}
}

// List handles GET /api/v1/groups
func (h *GroupHandler) List(w http.ResponseWriter, r *http.Request) {
	groups, err := h.groupRepo.List(r.Context())
	if err != nil {
		InternalError(w, r, "failed to list groups")
		return
	}

	OK(w, r, groups)
}

// Get handles GET /api/v1/groups/{id}
func (h *GroupHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid group id")
		return
	}

	group, err := h.groupRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "group")
		return
	}

	OK(w, r, group)
}

// Create handles POST /api/v1/groups
func (h *GroupHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	id, err := h.groupRepo.Create(r.Context(), req.Name, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create group")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/groups/{id}
func (h *GroupHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid group id")
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

	if err := h.groupRepo.Update(r.Context(), id, req.Name, req.Description); err != nil {
		NotFound(w, r, "group")
		return
	}

	OK(w, r, map[string]string{"message": "group updated"})
}

// Delete handles DELETE /api/v1/groups/{id}
func (h *GroupHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid group id")
		return
	}

	if err := h.groupRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "group")
		return
	}

	OK(w, r, map[string]string{"message": "group deleted"})
}

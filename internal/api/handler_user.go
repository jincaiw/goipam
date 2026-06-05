package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/auth"
	"github.com/jincaiw/goipam/internal/rbac"
	"github.com/jincaiw/goipam/internal/repository"
)

// UserHandler handles user management API endpoints.
type UserHandler struct {
	userRepo *repository.UserRepository
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// List handles GET /api/v1/users
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	users, err := h.userRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list users")
		return
	}

	total, err := h.userRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count users")
		return
	}

	OKPaged(w, r, users, total, page, pageSize)
}

// Get handles GET /api/v1/users/{id}
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid user id")
		return
	}

	user, err := h.userRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "user")
		return
	}

	OK(w, r, user)
}

// Create handles POST /api/v1/users
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" {
		BadRequest(w, r, errcode.ValidationError, "username and password are required")
		return
	}

	hash, err := auth.HashPassword(req.Password, auth.DefaultParams)
	if err != nil {
		InternalError(w, r, "failed to hash password")
		return
	}

	id, err := h.userRepo.Create(r.Context(), req.Username, hash, req.DisplayName, req.Email, "active")
	if err != nil {
		Error(w, r, http.StatusConflict, errcode.ResourceConflict, "username already exists")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/users/{id}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid user id")
		return
	}

	var req struct {
		DisplayName string `json:"display_name"`
		Email       string `json:"email"`
		Status      string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if err := h.userRepo.Update(r.Context(), id, req.DisplayName, req.Email, req.Status); err != nil {
		NotFound(w, r, "user")
		return
	}

	OK(w, r, map[string]string{"message": "user updated"})
}

// Delete handles DELETE /api/v1/users/{id}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid user id")
		return
	}

	// Prevent self-deletion
	if id == rbac.GetUserID(r.Context()) {
		BadRequest(w, r, errcode.ValidationError, "cannot delete yourself")
		return
	}

	if err := h.userRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "user")
		return
	}

	OK(w, r, map[string]string{"message": "user deleted"})
}

// parsePagination extracts page and page_size from query parameters.
func parsePagination(r *http.Request) (page, pageSize int) {
	page = 1
	pageSize = 20

	if v := r.URL.Query().Get("page"); v != "" {
		if p, err := strconv.Atoi(v); err == nil && p > 0 {
			page = p
		}
	}
	if v := r.URL.Query().Get("page_size"); v != "" {
		if ps, err := strconv.Atoi(v); err == nil && ps > 0 && ps <= 100 {
			pageSize = ps
		}
	}
	return
}

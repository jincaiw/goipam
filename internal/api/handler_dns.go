package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// NameserverHandler handles nameserver API endpoints.
type NameserverHandler struct {
	nameserverRepo *repository.NameserverRepository
}

// NewNameserverHandler creates a new NameserverHandler.
func NewNameserverHandler(nameserverRepo *repository.NameserverRepository) *NameserverHandler {
	return &NameserverHandler{nameserverRepo: nameserverRepo}
}

// List handles GET /api/v1/nameservers
func (h *NameserverHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	nameservers, err := h.nameserverRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list nameservers")
		return
	}

	total, err := h.nameserverRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count nameservers")
		return
	}

	OKPaged(w, r, nameservers, total, page, pageSize)
}

// Get handles GET /api/v1/nameservers/{id}
func (h *NameserverHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nameserver id")
		return
	}

	ns, err := h.nameserverRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "nameserver")
		return
	}

	OK(w, r, ns)
}

// Create handles POST /api/v1/nameservers
func (h *NameserverHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Server      string `json:"server"`
		Port        int32  `json:"port"`
		TimeoutMs   int32  `json:"timeout_ms"`
		Enabled     bool   `json:"enabled"`
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
	if req.Server == "" {
		BadRequest(w, r, errcode.ValidationError, "server is required")
		return
	}
	if req.Type == "" {
		req.Type = "udp"
	}
	if req.Port == 0 {
		req.Port = 53
	}
	if req.TimeoutMs == 0 {
		req.TimeoutMs = 2000
	}

	id, err := h.nameserverRepo.Create(
		r.Context(), req.Name, req.Type, req.Server,
		req.Port, req.TimeoutMs, req.Enabled, req.Description,
	)
	if err != nil {
		Error(w, r, http.StatusConflict, errcode.ResourceConflict, "nameserver name already exists")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/nameservers/{id}
func (h *NameserverHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nameserver id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Server      string `json:"server"`
		Port        int32  `json:"port"`
		TimeoutMs   int32  `json:"timeout_ms"`
		Enabled     bool   `json:"enabled"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if err := h.nameserverRepo.Update(
		r.Context(), id, req.Name, req.Type, req.Server,
		req.Port, req.TimeoutMs, req.Enabled, req.Description,
	); err != nil {
		NotFound(w, r, "nameserver")
		return
	}

	OK(w, r, map[string]string{"message": "nameserver updated"})
}

// Delete handles DELETE /api/v1/nameservers/{id}
func (h *NameserverHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid nameserver id")
		return
	}

	if err := h.nameserverRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "nameserver")
		return
	}

	OK(w, r, map[string]string{"message": "nameserver deleted"})
}

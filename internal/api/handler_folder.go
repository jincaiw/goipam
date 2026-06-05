package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// FolderHandler handles folder API endpoints.
type FolderHandler struct {
	folderRepo *repository.FolderRepository
}

// NewFolderHandler creates a new FolderHandler.
func NewFolderHandler(folderRepo *repository.FolderRepository) *FolderHandler {
	return &FolderHandler{folderRepo: folderRepo}
}

// List handles GET /api/v1/folders
func (h *FolderHandler) List(w http.ResponseWriter, r *http.Request) {
	sectionIDStr := r.URL.Query().Get("section_id")
	if sectionIDStr == "" {
		BadRequest(w, r, errcode.ValidationError, "section_id query parameter is required")
		return
	}

	sectionID, err := strconv.ParseInt(sectionIDStr, 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section_id")
		return
	}

	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	folders, err := h.folderRepo.List(r.Context(), sectionID, int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list folders")
		return
	}

	total, err := h.folderRepo.Count(r.Context(), sectionID)
	if err != nil {
		InternalError(w, r, "failed to count folders")
		return
	}

	OKPaged(w, r, folders, total, page, pageSize)
}

// Get handles GET /api/v1/folders/{id}
func (h *FolderHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid folder id")
		return
	}

	folder, err := h.folderRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "folder")
		return
	}

	OK(w, r, folder)
}

// Create handles POST /api/v1/folders
func (h *FolderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SectionID   int64  `json:"section_id"`
		ParentID    *int64 `json:"parent_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		SortOrder   int32  `json:"sort_order"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	if req.SectionID == 0 {
		BadRequest(w, r, errcode.ValidationError, "section_id is required")
		return
	}

	id, err := h.folderRepo.Create(r.Context(), req.SectionID, req.ParentID, req.Name, req.Description, req.SortOrder)
	if err != nil {
		InternalError(w, r, "failed to create folder")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/folders/{id}
func (h *FolderHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid folder id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		SortOrder   int32  `json:"sort_order"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	if err := h.folderRepo.Update(r.Context(), id, req.Name, req.Description, req.SortOrder); err != nil {
		NotFound(w, r, "folder")
		return
	}

	OK(w, r, map[string]string{"message": "folder updated"})
}

// Delete handles DELETE /api/v1/folders/{id}
func (h *FolderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid folder id")
		return
	}

	// Check for child folders
	children, err := h.folderRepo.GetSubFolders(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child folders")
		return
	}
	if len(children) > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "folder has child folders")
		return
	}

	if err := h.folderRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "folder")
		return
	}

	OK(w, r, map[string]string{"message": "folder deleted"})
}

package api

import (
	"encoding/json"
	"net/http"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// SettingsHandler handles settings API endpoints.
type SettingsHandler struct {
	repo *repository.SettingsRepository
}

// NewSettingsHandler creates a new SettingsHandler.
func NewSettingsHandler(repo *repository.SettingsRepository) *SettingsHandler {
	return &SettingsHandler{repo: repo}
}

// Get handles GET /api/v1/settings/{category}
func (h *SettingsHandler) Get(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	if category == "" {
		BadRequest(w, r, errcode.ValidationError, "category is required")
		return
	}

	settings, err := h.repo.GetByCategory(r.Context(), category)
	if err != nil {
		InternalError(w, r, "failed to get settings")
		return
	}

	OK(w, r, settings)
}

// Save handles PUT /api/v1/settings/{category}
func (h *SettingsHandler) Save(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	if category == "" {
		BadRequest(w, r, errcode.ValidationError, "category is required")
		return
	}

	var settings map[string]string
	if err := json.NewDecoder(r.Body).Decode(&settings); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if err := h.repo.UpsertBatch(r.Context(), category, settings); err != nil {
		InternalError(w, r, "failed to save settings")
		return
	}

	OK(w, r, map[string]string{"message": "settings saved"})
}

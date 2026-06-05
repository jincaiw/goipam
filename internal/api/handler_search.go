package api

import (
	"net/http"

	"github.com/jincaiw/goipam/internal/repository"
)

// SearchHandler handles search API endpoints.
type SearchHandler struct {
	searchRepo *repository.SearchRepository
}

// NewSearchHandler creates a new SearchHandler.
func NewSearchHandler(searchRepo *repository.SearchRepository) *SearchHandler {
	return &SearchHandler{searchRepo: searchRepo}
}

// Search handles GET /api/v1/search?q=keyword
func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		BadRequest(w, r, "VALIDATION_ERROR", "search keyword 'q' is required")
		return
	}

	ctx := r.Context()

	// Search addresses (limit 50)
	addresses, err := h.searchRepo.SearchAddresses(ctx, keyword, 50, 0)
	if err != nil {
		InternalError(w, r, "failed to search addresses")
		return
	}

	// Search subnets (limit 20)
	subnets, err := h.searchRepo.SearchSubnets(ctx, keyword, 20)
	if err != nil {
		InternalError(w, r, "failed to search subnets")
		return
	}

	// Search devices (limit 20)
	devices, err := h.searchRepo.SearchDevices(ctx, keyword, 20)
	if err != nil {
		InternalError(w, r, "failed to search devices")
		return
	}

	OK(w, r, repository.SearchResult{
		Addresses: addresses,
		Subnets:   subnets,
		Devices:   devices,
	})
}

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/repository"
)

// SectionHandler handles section API endpoints.
type SectionHandler struct {
	sectionRepo *repository.SectionRepository
	folderRepo  *repository.FolderRepository
	subnetRepo  *repository.SubnetRepository
	addressRepo *repository.AddressRepository
}

// NewSectionHandler creates a new SectionHandler.
func NewSectionHandler(sectionRepo *repository.SectionRepository, folderRepo *repository.FolderRepository, subnetRepo *repository.SubnetRepository, addressRepo *repository.AddressRepository) *SectionHandler {
	return &SectionHandler{
		sectionRepo: sectionRepo,
		folderRepo:  folderRepo,
		subnetRepo:  subnetRepo,
		addressRepo: addressRepo,
	}
}

// List handles GET /api/v1/sections
func (h *SectionHandler) List(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	sections, err := h.sectionRepo.List(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list sections")
		return
	}

	total, err := h.sectionRepo.Count(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count sections")
		return
	}

	OKPaged(w, r, sections, total, page, pageSize)
}

// Get handles GET /api/v1/sections/{id}
func (h *SectionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section id")
		return
	}

	section, err := h.sectionRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "section")
		return
	}

	OK(w, r, section)
}

// Create handles POST /api/v1/sections
func (h *SectionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    *int64 `json:"parent_id"`
		StrictMode  bool   `json:"strict_mode"`
		ShowVlan    bool   `json:"show_vlan"`
		ShowVrf     bool   `json:"show_vrf"`
		SubnetOrder string `json:"subnet_order"`
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

	if req.SubnetOrder == "" {
		req.SubnetOrder = "subnet"
	}

	id, err := h.sectionRepo.Create(r.Context(), req.Name, req.Description, req.ParentID, req.StrictMode, req.ShowVlan, req.ShowVrf, req.SubnetOrder, req.SortOrder)
	if err != nil {
		InternalError(w, r, "failed to create section")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/sections/{id}
func (h *SectionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ParentID    *int64 `json:"parent_id"`
		StrictMode  bool   `json:"strict_mode"`
		ShowVlan    bool   `json:"show_vlan"`
		ShowVrf     bool   `json:"show_vrf"`
		SubnetOrder string `json:"subnet_order"`
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

	if err := h.sectionRepo.Update(r.Context(), id, req.Name, req.Description, req.ParentID, req.StrictMode, req.ShowVlan, req.ShowVrf, req.SubnetOrder, req.SortOrder); err != nil {
		NotFound(w, r, "section")
		return
	}

	OK(w, r, map[string]string{"message": "section updated"})
}

// Delete handles DELETE /api/v1/sections/{id}
func (h *SectionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section id")
		return
	}

	// Check for child sections
	children, err := h.sectionRepo.GetSubSections(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child sections")
		return
	}
	if len(children) > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "section has child sections")
		return
	}

	// Check for child folders
	folderCount, err := h.folderRepo.Count(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child folders")
		return
	}
	if folderCount > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "section has child folders")
		return
	}

	// Check for child subnets
	subnetCount, err := h.subnetRepo.Count(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child subnets")
		return
	}
	if subnetCount > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "section has child subnets")
		return
	}

	if err := h.sectionRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "section")
		return
	}

	OK(w, r, map[string]string{"message": "section deleted"})
}

// Tree handles GET /api/v1/sections/{id}/tree
func (h *SectionHandler) Tree(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section id")
		return
	}

	section, err := h.sectionRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "section")
		return
	}

	folders, err := h.folderRepo.List(r.Context(), id, 1000, 0)
	if err != nil {
		InternalError(w, r, "failed to list folders")
		return
	}

	subnets, err := h.subnetRepo.List(r.Context(), id, 1000, 0)
	if err != nil {
		InternalError(w, r, "failed to list subnets")
		return
	}

	OK(w, r, map[string]interface{}{
		"section": section,
		"folders": folders,
		"subnets": subnets,
	})
}

// Stats handles GET /api/v1/sections/{id}/stats
func (h *SectionHandler) Stats(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid section id")
		return
	}

	_, err = h.sectionRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "section")
		return
	}

	subnetCount, err := h.subnetRepo.Count(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to count subnets")
		return
	}

	subnets, err := h.subnetRepo.List(r.Context(), id, 1000, 0)
	if err != nil {
		InternalError(w, r, "failed to list subnets")
		return
	}

	var totalAddresses int64
	var usedAddresses int64
	for _, s := range subnets {
		info, err := ipam.ParseCIDR(s.Cidr)
		if err != nil {
			continue
		}
		totalAddresses += info.UsableAddresses

		addrCount, err := h.addressRepo.Count(r.Context(), s.ID)
		if err != nil {
			continue
		}
		usedAddresses += addrCount
	}

	OK(w, r, map[string]interface{}{
		"subnet_count":    subnetCount,
		"address_count":   usedAddresses,
		"total_addresses": totalAddresses,
		"usage_percent":   usagePercent(totalAddresses, usedAddresses),
	})
}

func usagePercent(total, used int64) float64 {
	if total == 0 {
		return 0
	}
	return float64(used) / float64(total) * 100
}

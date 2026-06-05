package api

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/domain/ipam"
	"github.com/jincaiw/goipam/internal/repository"
)

// SubnetHandler handles subnet API endpoints.
type SubnetHandler struct {
	subnetRepo  *repository.SubnetRepository
	addressRepo *repository.AddressRepository
}

// NewSubnetHandler creates a new SubnetHandler.
func NewSubnetHandler(subnetRepo *repository.SubnetRepository, addressRepo *repository.AddressRepository) *SubnetHandler {
	return &SubnetHandler{
		subnetRepo:  subnetRepo,
		addressRepo: addressRepo,
	}
}

// List handles GET /api/v1/subnets
func (h *SubnetHandler) List(w http.ResponseWriter, r *http.Request) {
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

	subnets, err := h.subnetRepo.List(r.Context(), sectionID, int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list subnets")
		return
	}

	total, err := h.subnetRepo.Count(r.Context(), sectionID)
	if err != nil {
		InternalError(w, r, "failed to count subnets")
		return
	}

	OKPaged(w, r, subnets, total, page, pageSize)
}

// Get handles GET /api/v1/subnets/{id}
func (h *SubnetHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	OK(w, r, subnet)
}

// Create handles POST /api/v1/subnets
func (h *SubnetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SectionID        int64  `json:"section_id"`
		ParentID         *int64 `json:"parent_id"`
		Cidr             string `json:"cidr"`
		Name             string `json:"name"`
		Description      string `json:"description"`
		VlanID           *int64 `json:"vlan_id"`
		VrfID            *int64 `json:"vrf_id"`
		LocationID       *int64 `json:"location_id"`
		NameserverID     *int64 `json:"nameserver_id"`
		ScanEnabled      bool   `json:"scan_enabled"`
		DiscoverEnabled  bool   `json:"discover_enabled"`
		ThresholdPercent int32  `json:"threshold_percent"`
		AllowRequest     bool   `json:"allow_request"`
		Status           string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Cidr == "" {
		BadRequest(w, r, errcode.ValidationError, "cidr is required")
		return
	}

	if req.SectionID == 0 {
		BadRequest(w, r, errcode.ValidationError, "section_id is required")
		return
	}

	// Validate CIDR
	if err := ipam.ValidateCIDR(req.Cidr); err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, err.Error())
		return
	}

	// Parse CIDR for metadata
	info, err := ipam.ParseCIDR(req.Cidr)
	if err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, err.Error())
		return
	}

	// Convert network address to binary for storage
	// info.Network is like "192.168.1.0/24", we need just the IP part
	_, ipNet, _ := net.ParseCIDR(info.Network)
	networkBytes := ipNet.IP
	if v4 := networkBytes.To4(); v4 != nil {
		networkBytes = v4
	}

	// Check for overlapping subnets - fetch candidates and check with CIDR comparison
	candidates, err := h.subnetRepo.FindOverlappingSubnets(r.Context(), req.SectionID, req.VrfID, 0)
	if err != nil {
		InternalError(w, r, "failed to check overlapping subnets")
		return
	}
	var overlapCIDRs []string
	for _, s := range candidates {
		hasOverlap, checkErr := ipam.CheckOverlap(info.Network, s.Cidr)
		if checkErr != nil {
			continue
		}
		if hasOverlap {
			overlapCIDRs = append(overlapCIDRs, s.Cidr)
		}
	}
	if len(overlapCIDRs) > 0 {
		Error(w, r, http.StatusConflict, errcode.SubnetOverlap, "subnet overlaps with existing subnets: "+joinCIDRs(overlapCIDRs))
		return
	}

	if req.Status == "" {
		req.Status = "active"
	}

	id, err := h.subnetRepo.Create(
		r.Context(), req.SectionID, req.ParentID, info.Network,
		uint8(info.IPVersion), networkBytes, uint8(info.PrefixLen),
		req.Name, req.Description, req.VlanID, req.VrfID,
		req.LocationID, req.NameserverID, req.ScanEnabled,
		req.DiscoverEnabled, req.ThresholdPercent, req.AllowRequest, req.Status,
	)
	if err != nil {
		InternalError(w, r, "failed to create subnet")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// Update handles PUT /api/v1/subnets/{id}
func (h *SubnetHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	var req struct {
		Name             string `json:"name"`
		Description      string `json:"description"`
		VlanID           *int64 `json:"vlan_id"`
		VrfID            *int64 `json:"vrf_id"`
		LocationID       *int64 `json:"location_id"`
		NameserverID     *int64 `json:"nameserver_id"`
		ScanEnabled      bool   `json:"scan_enabled"`
		DiscoverEnabled  bool   `json:"discover_enabled"`
		ThresholdPercent int32  `json:"threshold_percent"`
		AllowRequest     bool   `json:"allow_request"`
		Status           string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if err := h.subnetRepo.Update(r.Context(), id, req.Name, req.Description, req.VlanID, req.VrfID, req.LocationID, req.NameserverID, req.ScanEnabled, req.DiscoverEnabled, req.ThresholdPercent, req.AllowRequest, req.Status); err != nil {
		NotFound(w, r, "subnet")
		return
	}

	OK(w, r, map[string]string{"message": "subnet updated"})
}

// Delete handles DELETE /api/v1/subnets/{id}
func (h *SubnetHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	// Check for child subnets
	children, err := h.subnetRepo.GetChildSubnets(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child subnets")
		return
	}
	if len(children) > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "subnet has child subnets")
		return
	}

	// Check for addresses
	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	addrCount, err := h.addressRepo.Count(r.Context(), subnet.ID)
	if err != nil {
		InternalError(w, r, "failed to count addresses")
		return
	}
	if addrCount > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "subnet has addresses")
		return
	}

	if err := h.subnetRepo.Delete(r.Context(), id); err != nil {
		NotFound(w, r, "subnet")
		return
	}

	OK(w, r, map[string]string{"message": "subnet deleted"})
}

// Children handles GET /api/v1/subnets/{id}/children
func (h *SubnetHandler) Children(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	children, err := h.subnetRepo.GetChildSubnets(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to get child subnets")
		return
	}

	OK(w, r, children)
}

// Addresses handles GET /api/v1/subnets/{id}/addresses
func (h *SubnetHandler) Addresses(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	addresses, err := h.addressRepo.List(r.Context(), id, int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list addresses")
		return
	}

	total, err := h.addressRepo.Count(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to count addresses")
		return
	}

	OKPaged(w, r, addresses, total, page, pageSize)
}

// FirstFreeIP handles GET /api/v1/subnets/{id}/first-free-ip
func (h *SubnetHandler) FirstFreeIP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	existingIPs, err := h.addressRepo.ListIPsBySubnet(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to list existing IPs")
		return
	}

	ip, err := ipam.GetFirstAvailableIP(subnet.Cidr, existingIPs)
	if err != nil {
		Error(w, r, http.StatusNotFound, errcode.NoAvailableIP, err.Error())
		return
	}

	OK(w, r, map[string]string{"ip": ip})
}

// AvailableRanges handles GET /api/v1/subnets/{id}/available-ranges
func (h *SubnetHandler) AvailableRanges(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	existingIPs, err := h.addressRepo.ListIPsBySubnet(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to list existing IPs")
		return
	}

	ranges, err := ipam.GetAvailableRanges(subnet.Cidr, existingIPs)
	if err != nil {
		InternalError(w, r, "failed to get available ranges")
		return
	}

	OK(w, r, ranges)
}

// Stats handles GET /api/v1/subnets/{id}/stats
func (h *SubnetHandler) Stats(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	info, err := ipam.ParseCIDR(subnet.Cidr)
	if err != nil {
		InternalError(w, r, "failed to parse subnet CIDR")
		return
	}

	usedCount, err := h.addressRepo.Count(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to count addresses")
		return
	}

	OK(w, r, map[string]interface{}{
		"cidr":             subnet.Cidr,
		"total_addresses":  info.TotalAddresses,
		"usable_addresses": info.UsableAddresses,
		"used_addresses":   usedCount,
		"free_addresses":   info.UsableAddresses - usedCount,
		"usage_percent":    usagePercent(info.UsableAddresses, usedCount),
	})
}

// CheckOverlap handles POST /api/v1/subnets/check-overlap
func (h *SubnetHandler) CheckOverlap(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SectionID int64  `json:"section_id"`
		Cidr      string `json:"cidr"`
		VrfID     *int64 `json:"vrf_id"`
		ExcludeID int64  `json:"exclude_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Cidr == "" {
		BadRequest(w, r, errcode.ValidationError, "cidr is required")
		return
	}

	if req.SectionID == 0 {
		BadRequest(w, r, errcode.ValidationError, "section_id is required")
		return
	}

	// Validate CIDR
	if err := ipam.ValidateCIDR(req.Cidr); err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, err.Error())
		return
	}

	info, err := ipam.ParseCIDR(req.Cidr)
	if err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, err.Error())
		return
	}

	candidates, err := h.subnetRepo.FindOverlappingSubnets(r.Context(), req.SectionID, req.VrfID, req.ExcludeID)
	if err != nil {
		InternalError(w, r, "failed to check overlapping subnets")
		return
	}

	var overlapping []map[string]interface{}
	for _, s := range candidates {
		hasOverlap, checkErr := ipam.CheckOverlap(info.Network, s.Cidr)
		if checkErr != nil {
			continue
		}
		if hasOverlap {
			overlapping = append(overlapping, map[string]interface{}{
				"id":   s.ID,
				"cidr": s.Cidr,
				"name": s.Name,
			})
		}
	}

	OK(w, r, map[string]interface{}{
		"overlaps":    len(overlapping) > 0,
		"overlapping": overlapping,
	})
}

// Split handles POST /api/v1/subnets/{id}/split
func (h *SubnetHandler) Split(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid subnet id")
		return
	}

	subnet, err := h.subnetRepo.GetByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "subnet")
		return
	}

	var req struct {
		NamePrefix string `json:"name_prefix"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Body is optional, ignore decode errors
		req.NamePrefix = ""
	}

	halves, err := ipam.SplitSubnet(subnet.Cidr)
	if err != nil {
		BadRequest(w, r, errcode.SubnetInvalidCIDR, "cannot split subnet: "+err.Error())
		return
	}

	// Update the existing subnet to the first half
	firstHalf := halves[0]
	_, firstIPNet, _ := net.ParseCIDR(firstHalf)
	networkBytes := firstIPNet.IP
	if v4 := networkBytes.To4(); v4 != nil {
		networkBytes = v4
	}

	firstInfo, _ := ipam.ParseCIDR(firstHalf)
	if err := h.subnetRepo.UpdateCIDR(r.Context(), id, firstHalf, uint8(firstInfo.IPVersion), networkBytes, uint8(firstInfo.PrefixLen)); err != nil {
		InternalError(w, r, "failed to update subnet to first half")
		return
	}

	// Create a new subnet for the second half
	secondHalf := halves[1]
	_, secondIPNet, _ := net.ParseCIDR(secondHalf)
	secondNetworkBytes := secondIPNet.IP
	if v4 := secondNetworkBytes.To4(); v4 != nil {
		secondNetworkBytes = v4
	}

	secondInfo, _ := ipam.ParseCIDR(secondHalf)
	secondName := req.NamePrefix + " (2nd half)"
	if subnet.Name != "" {
		secondName = subnet.Name + " (2nd half)"
	}
	if req.NamePrefix != "" {
		secondName = req.NamePrefix + " (2nd half)"
	}

	secondID, err := h.subnetRepo.Create(
		r.Context(), subnet.SectionID, &id, secondHalf,
		uint8(secondInfo.IPVersion), secondNetworkBytes, uint8(secondInfo.PrefixLen),
		secondName, subnet.Description, subnet.VlanID, subnet.VrfID,
		subnet.LocationID, subnet.NameserverID, subnet.ScanEnabled,
		subnet.DiscoverEnabled, subnet.ThresholdPercent, subnet.AllowRequest, subnet.Status,
	)
	if err != nil {
		InternalError(w, r, "failed to create second half subnet")
		return
	}

	OK(w, r, map[string]interface{}{
		"first_half":  map[string]interface{}{"id": id, "cidr": firstHalf},
		"second_half": map[string]interface{}{"id": secondID, "cidr": secondHalf},
	})
}

func joinCIDRs(cidrs []string) string {
	result := cidrs[0]
	for _, c := range cidrs[1:] {
		result += ", " + c
	}
	return result
}

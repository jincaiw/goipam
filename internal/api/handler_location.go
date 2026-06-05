package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// LocationHandler handles location and rack API endpoints.
type LocationHandler struct {
	locationRepo *repository.LocationRepository
	deviceRepo   *repository.DeviceRepository
}

// NewLocationHandler creates a new LocationHandler.
func NewLocationHandler(locationRepo *repository.LocationRepository, deviceRepo *repository.DeviceRepository) *LocationHandler {
	return &LocationHandler{locationRepo: locationRepo, deviceRepo: deviceRepo}
}

// ListLocations handles GET /api/v1/locations
func (h *LocationHandler) ListLocations(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	locations, err := h.locationRepo.ListLocations(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list locations")
		return
	}

	total, err := h.locationRepo.CountLocations(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count locations")
		return
	}

	OKPaged(w, r, locations, total, page, pageSize)
}

// GetLocation handles GET /api/v1/locations/{id}
func (h *LocationHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid location id")
		return
	}

	location, err := h.locationRepo.GetLocationByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "location")
		return
	}

	OK(w, r, location)
}

// CreateLocation handles POST /api/v1/locations
func (h *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Address     string `json:"address"`
		Room        string `json:"room"`
		RowName     string `json:"row_name"`
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

	id, err := h.locationRepo.CreateLocation(r.Context(), req.Name, req.Address, req.Room, req.RowName, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create location")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// UpdateLocation handles PUT /api/v1/locations/{id}
func (h *LocationHandler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid location id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		Address     string `json:"address"`
		Room        string `json:"room"`
		RowName     string `json:"row_name"`
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

	if err := h.locationRepo.UpdateLocation(r.Context(), id, req.Name, req.Address, req.Room, req.RowName, req.Description); err != nil {
		NotFound(w, r, "location")
		return
	}

	OK(w, r, map[string]string{"message": "location updated"})
}

// DeleteLocation handles DELETE /api/v1/locations/{id}
func (h *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid location id")
		return
	}

	// Check for child racks
	rackCount, err := h.locationRepo.CountRacks(r.Context(), id)
	if err != nil {
		InternalError(w, r, "failed to check child racks")
		return
	}
	if rackCount > 0 {
		Error(w, r, http.StatusConflict, errcode.ResourceHasChildren, "location has child racks")
		return
	}

	if err := h.locationRepo.DeleteLocation(r.Context(), id); err != nil {
		NotFound(w, r, "location")
		return
	}

	OK(w, r, map[string]string{"message": "location deleted"})
}

// ListRacks handles GET /api/v1/racks
func (h *LocationHandler) ListRacks(w http.ResponseWriter, r *http.Request) {
	locationID := int64(0)
	locationIDStr := r.URL.Query().Get("location_id")
	if locationIDStr != "" {
		var err error
		locationID, err = strconv.ParseInt(locationIDStr, 10, 64)
		if err != nil {
			BadRequest(w, r, errcode.ValidationError, "invalid location_id")
			return
		}
	}

	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	racks, err := h.locationRepo.ListRacks(r.Context(), locationID, int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list racks")
		return
	}

	total, err := h.locationRepo.CountRacks(r.Context(), locationID)
	if err != nil {
		InternalError(w, r, "failed to count racks")
		return
	}

	OKPaged(w, r, racks, total, page, pageSize)
}

// GetRack handles GET /api/v1/racks/{id}
func (h *LocationHandler) GetRack(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid rack id")
		return
	}

	rack, err := h.locationRepo.GetRackByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "rack")
		return
	}

	OK(w, r, rack)
}

// CreateRack handles POST /api/v1/racks
func (h *LocationHandler) CreateRack(w http.ResponseWriter, r *http.Request) {
	var req struct {
		LocationID  int64  `json:"location_id"`
		Name        string `json:"name"`
		SizeU       int32  `json:"size_u"`
		RowName     string `json:"row_name"`
		Position    string `json:"position"`
		Direction   string `json:"direction"`
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
	if req.LocationID == 0 {
		BadRequest(w, r, errcode.ValidationError, "location_id is required")
		return
	}
	if req.SizeU < 1 {
		BadRequest(w, r, errcode.ValidationError, "size_u must be positive")
		return
	}

	id, err := h.locationRepo.CreateRack(r.Context(), req.LocationID, req.Name, req.SizeU, req.RowName, req.Position, req.Direction, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create rack")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// UpdateRack handles PUT /api/v1/racks/{id}
func (h *LocationHandler) UpdateRack(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid rack id")
		return
	}

	var req struct {
		Name        string `json:"name"`
		SizeU       int32  `json:"size_u"`
		RowName     string `json:"row_name"`
		Position    string `json:"position"`
		Direction   string `json:"direction"`
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
	if req.SizeU < 1 {
		BadRequest(w, r, errcode.ValidationError, "size_u must be positive")
		return
	}

	if err := h.locationRepo.UpdateRack(r.Context(), id, req.Name, req.SizeU, req.RowName, req.Position, req.Direction, req.Description); err != nil {
		NotFound(w, r, "rack")
		return
	}

	OK(w, r, map[string]string{"message": "rack updated"})
}

// DeleteRack handles DELETE /api/v1/racks/{id}
func (h *LocationHandler) DeleteRack(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid rack id")
		return
	}

	if err := h.locationRepo.DeleteRack(r.Context(), id); err != nil {
		NotFound(w, r, "rack")
		return
	}

	OK(w, r, map[string]string{"message": "rack deleted"})
}

// MountDevice handles POST /api/v1/racks/{id}/mount-device
func (h *LocationHandler) MountDevice(w http.ResponseWriter, r *http.Request) {
	rackID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid rack id")
		return
	}

	var req struct {
		DeviceID int64 `json:"device_id"`
		Position int32 `json:"position"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.DeviceID == 0 {
		BadRequest(w, r, errcode.ValidationError, "device_id is required")
		return
	}

	// Verify rack exists
	if _, err := h.locationRepo.GetRackByID(r.Context(), rackID); err != nil {
		NotFound(w, r, "rack")
		return
	}

	// Verify device exists and get current state
	device, err := h.deviceRepo.GetDeviceByID(r.Context(), req.DeviceID)
	if err != nil {
		NotFound(w, r, "device")
		return
	}

	// Update device with rack assignment
	if err := h.deviceRepo.UpdateDevice(r.Context(), device.ID, device.Name, device.Hostname, device.MgmtIP,
		device.TypeID, device.LocationID, &rackID, &req.Position, device.RackSizeU,
		device.Manufacturer, device.Model, device.SerialNumber, device.Description); err != nil {
		InternalError(w, r, "failed to mount device to rack")
		return
	}

	OK(w, r, map[string]string{"message": "device mounted to rack"})
}

// UnmountDevice handles POST /api/v1/racks/{id}/unmount-device
func (h *LocationHandler) UnmountDevice(w http.ResponseWriter, r *http.Request) {
	rackID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid rack id")
		return
	}

	var req struct {
		DeviceID int64 `json:"device_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.DeviceID == 0 {
		BadRequest(w, r, errcode.ValidationError, "device_id is required")
		return
	}

	// Verify rack exists
	if _, err := h.locationRepo.GetRackByID(r.Context(), rackID); err != nil {
		NotFound(w, r, "rack")
		return
	}

	// Verify device exists and get current state
	device, err := h.deviceRepo.GetDeviceByID(r.Context(), req.DeviceID)
	if err != nil {
		NotFound(w, r, "device")
		return
	}

	// Update device to remove rack assignment
	if err := h.deviceRepo.UpdateDevice(r.Context(), device.ID, device.Name, device.Hostname, device.MgmtIP,
		device.TypeID, device.LocationID, nil, nil, device.RackSizeU,
		device.Manufacturer, device.Model, device.SerialNumber, device.Description); err != nil {
		InternalError(w, r, "failed to unmount device from rack")
		return
	}

	OK(w, r, map[string]string{"message": "device unmounted from rack"})
}

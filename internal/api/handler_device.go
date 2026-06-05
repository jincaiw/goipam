package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/repository"
)

// DeviceHandler handles device and device type API endpoints.
type DeviceHandler struct {
	deviceRepo *repository.DeviceRepository
}

// NewDeviceHandler creates a new DeviceHandler.
func NewDeviceHandler(deviceRepo *repository.DeviceRepository) *DeviceHandler {
	return &DeviceHandler{deviceRepo: deviceRepo}
}

// ListDevices handles GET /api/v1/devices
func (h *DeviceHandler) ListDevices(w http.ResponseWriter, r *http.Request) {
	page, pageSize := parsePagination(r)
	offset := int32((page - 1) * pageSize)

	devices, err := h.deviceRepo.ListDevices(r.Context(), int32(pageSize), offset)
	if err != nil {
		InternalError(w, r, "failed to list devices")
		return
	}

	total, err := h.deviceRepo.CountDevices(r.Context())
	if err != nil {
		InternalError(w, r, "failed to count devices")
		return
	}

	OKPaged(w, r, devices, total, page, pageSize)
}

// GetDevice handles GET /api/v1/devices/{id}
func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid device id")
		return
	}

	device, err := h.deviceRepo.GetDeviceByID(r.Context(), id)
	if err != nil {
		NotFound(w, r, "device")
		return
	}

	OK(w, r, device)
}

// CreateDevice handles POST /api/v1/devices
func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name         string `json:"name"`
		Hostname     string `json:"hostname"`
		MgmtIP       string `json:"mgmt_ip"`
		TypeID       *int64 `json:"type_id"`
		LocationID   *int64 `json:"location_id"`
		RackID       *int64 `json:"rack_id"`
		RackStartU   *int32 `json:"rack_start_u"`
		RackSizeU    *int32 `json:"rack_size_u"`
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		SerialNumber string `json:"serial_number"`
		Description  string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	id, err := h.deviceRepo.CreateDevice(r.Context(), req.Name, req.Hostname, req.MgmtIP, req.TypeID, req.LocationID, req.RackID, req.RackStartU, req.RackSizeU, req.Manufacturer, req.Model, req.SerialNumber, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create device")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

// UpdateDevice handles PUT /api/v1/devices/{id}
func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid device id")
		return
	}

	var req struct {
		Name         string `json:"name"`
		Hostname     string `json:"hostname"`
		MgmtIP       string `json:"mgmt_ip"`
		TypeID       *int64 `json:"type_id"`
		LocationID   *int64 `json:"location_id"`
		RackID       *int64 `json:"rack_id"`
		RackStartU   *int32 `json:"rack_start_u"`
		RackSizeU    *int32 `json:"rack_size_u"`
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		SerialNumber string `json:"serial_number"`
		Description  string `json:"description"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid request body")
		return
	}

	if req.Name == "" {
		BadRequest(w, r, errcode.ValidationError, "name is required")
		return
	}

	if err := h.deviceRepo.UpdateDevice(r.Context(), id, req.Name, req.Hostname, req.MgmtIP, req.TypeID, req.LocationID, req.RackID, req.RackStartU, req.RackSizeU, req.Manufacturer, req.Model, req.SerialNumber, req.Description); err != nil {
		NotFound(w, r, "device")
		return
	}

	OK(w, r, map[string]string{"message": "device updated"})
}

// DeleteDevice handles DELETE /api/v1/devices/{id}
func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		BadRequest(w, r, errcode.ValidationError, "invalid device id")
		return
	}

	if err := h.deviceRepo.DeleteDevice(r.Context(), id); err != nil {
		NotFound(w, r, "device")
		return
	}

	OK(w, r, map[string]string{"message": "device deleted"})
}

// ListDeviceTypes handles GET /api/v1/device-types
func (h *DeviceHandler) ListDeviceTypes(w http.ResponseWriter, r *http.Request) {
	types, err := h.deviceRepo.ListDeviceTypes(r.Context())
	if err != nil {
		InternalError(w, r, "failed to list device types")
		return
	}

	OK(w, r, types)
}

// CreateDeviceType handles POST /api/v1/device-types
func (h *DeviceHandler) CreateDeviceType(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Icon        string `json:"icon"`
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

	id, err := h.deviceRepo.CreateDeviceType(r.Context(), req.Name, req.Icon, req.Description)
	if err != nil {
		InternalError(w, r, "failed to create device type")
		return
	}

	Created(w, r, map[string]int64{"id": id})
}

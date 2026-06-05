package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Device represents a device domain object.
type Device struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Hostname     string    `json:"hostname"`
	MgmtIP       string    `json:"mgmt_ip"`
	TypeID       *int64    `json:"type_id"`
	LocationID   *int64    `json:"location_id"`
	RackID       *int64    `json:"rack_id"`
	RackStartU   *int32    `json:"rack_start_u"`
	RackSizeU    *int32    `json:"rack_size_u"`
	Manufacturer string    `json:"manufacturer"`
	Model        string    `json:"model"`
	SerialNumber string    `json:"serial_number"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// DeviceType represents a device type domain object.
type DeviceType struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DeviceRepository handles device and device type data access.
type DeviceRepository struct {
	q *sqlc.Queries
}

// NewDeviceRepository creates a new DeviceRepository.
func NewDeviceRepository(q *sqlc.Queries) *DeviceRepository {
	return &DeviceRepository{q: q}
}

// GetDeviceByID retrieves a device by ID.
func (r *DeviceRepository) GetDeviceByID(ctx context.Context, id int64) (*Device, error) {
	row, err := r.q.GetDeviceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get device by id: %w", err)
	}
	return sqlcDeviceToDevice(row), nil
}

// ListDevices retrieves a paginated list of devices.
func (r *DeviceRepository) ListDevices(ctx context.Context, limit, offset int32) ([]*Device, error) {
	rows, err := r.q.ListDevices(ctx, sqlc.ListDevicesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list devices: %w", err)
	}
	devices := make([]*Device, len(rows))
	for i, row := range rows {
		devices[i] = sqlcDeviceToDevice(row)
	}
	return devices, nil
}

// CountDevices returns the total number of devices.
func (r *DeviceRepository) CountDevices(ctx context.Context) (int64, error) {
	return r.q.CountDevices(ctx)
}

// CreateDevice creates a new device and returns the ID.
func (r *DeviceRepository) CreateDevice(ctx context.Context, name, hostname, mgmtIP string, typeID, locationID, rackID *int64, rackStartU, rackSizeU *int32, manufacturer, model, serialNumber, description string) (int64, error) {
	result, err := r.q.CreateDevice(ctx, sqlc.CreateDeviceParams{
		Name:         name,
		Hostname:     hostname,
		MgmtIp:       mgmtIP,
		TypeID:       NullInt64(typeID),
		LocationID:   NullInt64(locationID),
		RackID:       NullInt64(rackID),
		RackStartU:   nullInt32(rackStartU),
		RackSizeU:    nullInt32(rackSizeU),
		Manufacturer: manufacturer,
		Model:        model,
		SerialNumber: serialNumber,
		Description:  NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create device: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateDevice updates a device.
func (r *DeviceRepository) UpdateDevice(ctx context.Context, id int64, name, hostname, mgmtIP string, typeID, locationID, rackID *int64, rackStartU, rackSizeU *int32, manufacturer, model, serialNumber, description string) error {
	return r.q.UpdateDevice(ctx, sqlc.UpdateDeviceParams{
		Name:         name,
		Hostname:     hostname,
		MgmtIp:       mgmtIP,
		TypeID:       NullInt64(typeID),
		LocationID:   NullInt64(locationID),
		RackID:       NullInt64(rackID),
		RackStartU:   nullInt32(rackStartU),
		RackSizeU:    nullInt32(rackSizeU),
		Manufacturer: manufacturer,
		Model:        model,
		SerialNumber: serialNumber,
		Description:  NullString(&description),
		ID:           id,
	})
}

// DeleteDevice deletes a device.
func (r *DeviceRepository) DeleteDevice(ctx context.Context, id int64) error {
	return r.q.DeleteDevice(ctx, id)
}

// GetDeviceTypeByID retrieves a device type by ID.
func (r *DeviceRepository) GetDeviceTypeByID(ctx context.Context, id int64) (*DeviceType, error) {
	row, err := r.q.GetDeviceTypeByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get device type by id: %w", err)
	}
	return sqlcDeviceTypeToDeviceType(row), nil
}

// ListDeviceTypes retrieves all device types.
func (r *DeviceRepository) ListDeviceTypes(ctx context.Context) ([]*DeviceType, error) {
	rows, err := r.q.ListDeviceTypes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list device types: %w", err)
	}
	types := make([]*DeviceType, len(rows))
	for i, row := range rows {
		types[i] = sqlcDeviceTypeToDeviceType(row)
	}
	return types, nil
}

// CreateDeviceType creates a new device type and returns the ID.
func (r *DeviceRepository) CreateDeviceType(ctx context.Context, name, icon, description string) (int64, error) {
	result, err := r.q.CreateDeviceType(ctx, sqlc.CreateDeviceTypeParams{
		Name:        name,
		Icon:        icon,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create device type: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateDeviceType updates a device type.
func (r *DeviceRepository) UpdateDeviceType(ctx context.Context, id int64, name, icon, description string) error {
	return r.q.UpdateDeviceType(ctx, sqlc.UpdateDeviceTypeParams{
		Name:        name,
		Icon:        icon,
		Description: NullString(&description),
		ID:          id,
	})
}

// DeleteDeviceType deletes a device type.
func (r *DeviceRepository) DeleteDeviceType(ctx context.Context, id int64) error {
	return r.q.DeleteDeviceType(ctx, id)
}

func sqlcDeviceToDevice(row sqlc.Device) *Device {
	return &Device{
		ID:           row.ID,
		Name:         row.Name,
		Hostname:     row.Hostname,
		MgmtIP:       row.MgmtIp,
		TypeID:       nullInt64Ptr(row.TypeID),
		LocationID:   nullInt64Ptr(row.LocationID),
		RackID:       nullInt64Ptr(row.RackID),
		RackStartU:   nullInt32Ptr(row.RackStartU),
		RackSizeU:    nullInt32Ptr(row.RackSizeU),
		Manufacturer: row.Manufacturer,
		Model:        row.Model,
		SerialNumber: row.SerialNumber,
		Description:  row.Description.String,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

func sqlcDeviceTypeToDeviceType(row sqlc.DeviceType) *DeviceType {
	return &DeviceType{
		ID:          row.ID,
		Name:        row.Name,
		Icon:        row.Icon,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func nullInt32(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *v, Valid: true}
}

func nullInt32Ptr(v sql.NullInt32) *int32 {
	if v.Valid {
		return &v.Int32
	}
	return nil
}

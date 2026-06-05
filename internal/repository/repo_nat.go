package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Nat represents a NAT domain object.
type Nat struct {
	ID               int64     `json:"id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	InsideAddressID  *int64    `json:"inside_address_id"`
	OutsideAddressID *int64    `json:"outside_address_id"`
	InsideIP         string    `json:"inside_ip"`
	OutsideIP        string    `json:"outside_ip"`
	InsidePort       *int32    `json:"inside_port"`
	OutsidePort      *int32    `json:"outside_port"`
	Protocol         string    `json:"protocol"`
	DeviceID         *int64    `json:"device_id"`
	Description      string    `json:"description"`
	Enabled          bool      `json:"enabled"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// NatRepository handles NAT data access.
type NatRepository struct {
	q *sqlc.Queries
}

// NewNatRepository creates a new NatRepository.
func NewNatRepository(q *sqlc.Queries) *NatRepository {
	return &NatRepository{q: q}
}

// GetByID retrieves a NAT by ID.
func (r *NatRepository) GetByID(ctx context.Context, id int64) (*Nat, error) {
	row, err := r.q.GetNatByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get nat by id: %w", err)
	}
	return sqlcNatToNat(row), nil
}

// List retrieves a paginated list of NATs.
func (r *NatRepository) List(ctx context.Context, limit, offset int32) ([]*Nat, error) {
	rows, err := r.q.ListNats(ctx, sqlc.ListNatsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list nats: %w", err)
	}
	nats := make([]*Nat, len(rows))
	for i, row := range rows {
		nats[i] = sqlcNatToNat(row)
	}
	return nats, nil
}

// Count returns the total number of NATs.
func (r *NatRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountNats(ctx)
}

// Create creates a new NAT and returns the ID.
func (r *NatRepository) Create(ctx context.Context, name, natType string, insideAddressID, outsideAddressID *int64, insideIP, outsideIP string, insidePort, outsidePort *int32, protocol string, deviceID *int64, description string, enabled bool) (int64, error) {
	result, err := r.q.CreateNat(ctx, sqlc.CreateNatParams{
		Name:             name,
		Type:             natType,
		InsideAddressID:  NullInt64(insideAddressID),
		OutsideAddressID: NullInt64(outsideAddressID),
		InsideIp:         insideIP,
		OutsideIp:        outsideIP,
		InsidePort:       nullInt32(insidePort),
		OutsidePort:      nullInt32(outsidePort),
		Protocol:         protocol,
		DeviceID:         NullInt64(deviceID),
		Description:      NullString(&description),
		Enabled:          enabled,
	})
	if err != nil {
		return 0, fmt.Errorf("create nat: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a NAT.
func (r *NatRepository) Update(ctx context.Context, id int64, name, natType string, insideAddressID, outsideAddressID *int64, insideIP, outsideIP string, insidePort, outsidePort *int32, protocol string, deviceID *int64, description string, enabled bool) error {
	return r.q.UpdateNat(ctx, sqlc.UpdateNatParams{
		Name:             name,
		Type:             natType,
		InsideAddressID:  NullInt64(insideAddressID),
		OutsideAddressID: NullInt64(outsideAddressID),
		InsideIp:         insideIP,
		OutsideIp:        outsideIP,
		InsidePort:       nullInt32(insidePort),
		OutsidePort:      nullInt32(outsidePort),
		Protocol:         protocol,
		DeviceID:         NullInt64(deviceID),
		Description:      NullString(&description),
		Enabled:          enabled,
		ID:               id,
	})
}

// Delete deletes a NAT.
func (r *NatRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteNat(ctx, id)
}

func sqlcNatToNat(row sqlc.Nat) *Nat {
	return &Nat{
		ID:               row.ID,
		Name:             row.Name,
		Type:             row.Type,
		InsideAddressID:  nullInt64Ptr(row.InsideAddressID),
		OutsideAddressID: nullInt64Ptr(row.OutsideAddressID),
		InsideIP:         row.InsideIp,
		OutsideIP:        row.OutsideIp,
		InsidePort:       nullInt32Ptr(row.InsidePort),
		OutsidePort:      nullInt32Ptr(row.OutsidePort),
		Protocol:         row.Protocol,
		DeviceID:         nullInt64Ptr(row.DeviceID),
		Description:      row.Description.String,
		Enabled:          row.Enabled,
		CreatedAt:        row.CreatedAt,
		UpdatedAt:        row.UpdatedAt,
	}
}

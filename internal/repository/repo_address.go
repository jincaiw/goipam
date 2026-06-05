package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Address represents an address domain object.
type Address struct {
	ID           int64      `json:"id"`
	SubnetID     int64      `json:"subnet_id"`
	Ip           []byte     `json:"ip"`
	IpText       string     `json:"ip_text"`
	Hostname     string     `json:"hostname"`
	Mac          string     `json:"mac"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	TagID        *int64     `json:"tag_id"`
	IsGateway    bool       `json:"is_gateway"`
	DeviceID     *int64     `json:"device_id"`
	DevicePort   string     `json:"device_port"`
	Owner        string     `json:"owner"`
	LastSeenAt   *time.Time `json:"last_seen_at"`
	DnsCheckedAt *time.Time `json:"dns_checked_at"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// AddressRepository handles address data access.
type AddressRepository struct {
	q *sqlc.Queries
}

// NewAddressRepository creates a new AddressRepository.
func NewAddressRepository(q *sqlc.Queries) *AddressRepository {
	return &AddressRepository{q: q}
}

// GetByID retrieves an address by ID.
func (r *AddressRepository) GetByID(ctx context.Context, id int64) (*Address, error) {
	row, err := r.q.GetAddressByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get address by id: %w", err)
	}
	return sqlcAddressToAddress(row), nil
}

// List retrieves a paginated list of addresses in a subnet.
func (r *AddressRepository) List(ctx context.Context, subnetID int64, limit, offset int32) ([]*Address, error) {
	rows, err := r.q.ListAddresses(ctx, sqlc.ListAddressesParams{
		SubnetID: subnetID,
		Limit:    limit,
		Offset:   offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list addresses: %w", err)
	}
	addresses := make([]*Address, len(rows))
	for i, row := range rows {
		addresses[i] = sqlcAddressToAddress(row)
	}
	return addresses, nil
}

// Count returns the total number of addresses in a subnet.
func (r *AddressRepository) Count(ctx context.Context, subnetID int64) (int64, error) {
	return r.q.CountAddresses(ctx, subnetID)
}

// CountByStatus returns the number of addresses with a given status in a subnet.
func (r *AddressRepository) CountByStatus(ctx context.Context, subnetID int64, status string) (int64, error) {
	return r.q.CountAddressesByStatus(ctx, sqlc.CountAddressesByStatusParams{
		SubnetID: subnetID,
		Status:   status,
	})
}

// Create creates a new address and returns the ID.
func (r *AddressRepository) Create(ctx context.Context, subnetID int64, ip []byte, ipText, hostname, mac, description, status string, tagID *int64, isGateway bool, deviceID *int64, devicePort, owner string) (int64, error) {
	result, err := r.q.CreateAddress(ctx, sqlc.CreateAddressParams{
		SubnetID:    subnetID,
		Ip:          ip,
		IpText:      ipText,
		Hostname:    hostname,
		Mac:         mac,
		Description: NullString(&description),
		Status:      status,
		TagID:       NullInt64(tagID),
		IsGateway:   isGateway,
		DeviceID:    NullInt64(deviceID),
		DevicePort:  devicePort,
		Owner:       owner,
	})
	if err != nil {
		return 0, fmt.Errorf("create address: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates an address.
func (r *AddressRepository) Update(ctx context.Context, id int64, hostname, mac, description, status string, tagID *int64, isGateway bool, deviceID *int64, devicePort, owner string) error {
	return r.q.UpdateAddress(ctx, sqlc.UpdateAddressParams{
		Hostname:    hostname,
		Mac:         mac,
		Description: NullString(&description),
		Status:      status,
		TagID:       NullInt64(tagID),
		IsGateway:   isGateway,
		DeviceID:    NullInt64(deviceID),
		DevicePort:  devicePort,
		Owner:       owner,
		ID:          id,
	})
}

// Delete deletes an address.
func (r *AddressRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteAddress(ctx, id)
}

// GetByIPText retrieves an address by subnet ID and IP text.
func (r *AddressRepository) GetByIPText(ctx context.Context, subnetID int64, ipText string) (*Address, error) {
	row, err := r.q.GetAddressByIPText(ctx, sqlc.GetAddressByIPTextParams{
		SubnetID: subnetID,
		IpText:   ipText,
	})
	if err != nil {
		return nil, fmt.Errorf("get address by ip text: %w", err)
	}
	return sqlcAddressToAddress(row), nil
}

// ListIPsBySubnet returns all IP text strings for a subnet.
func (r *AddressRepository) ListIPsBySubnet(ctx context.Context, subnetID int64) ([]string, error) {
	return r.q.ListAddressIPsBySubnet(ctx, subnetID)
}

// Search searches addresses by hostname, IP text, MAC, or description.
func (r *AddressRepository) Search(ctx context.Context, keyword string, limit, offset int32) ([]*Address, error) {
	pattern := "%" + keyword + "%"
	rows, err := r.q.SearchAddresses(ctx, sqlc.SearchAddressesParams{
		Hostname:    pattern,
		IpText:      pattern,
		Mac:         pattern,
		Description: sql.NullString{String: pattern, Valid: true},
		Limit:       limit,
		Offset:      offset,
	})
	if err != nil {
		return nil, fmt.Errorf("search addresses: %w", err)
	}
	addresses := make([]*Address, len(rows))
	for i, row := range rows {
		addresses[i] = sqlcAddressToAddress(row)
	}
	return addresses, nil
}

func sqlcAddressToAddress(row sqlc.Address) *Address {
	return &Address{
		ID:           row.ID,
		SubnetID:     row.SubnetID,
		Ip:           row.Ip,
		IpText:       row.IpText,
		Hostname:     row.Hostname,
		Mac:          row.Mac,
		Description:  row.Description.String,
		Status:       row.Status,
		TagID:        nullInt64Ptr(row.TagID),
		IsGateway:    row.IsGateway,
		DeviceID:     nullInt64Ptr(row.DeviceID),
		DevicePort:   row.DevicePort,
		Owner:        row.Owner,
		LastSeenAt:   nullTimePtr(row.LastSeenAt),
		DnsCheckedAt: nullTimePtr(row.DnsCheckedAt),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
}

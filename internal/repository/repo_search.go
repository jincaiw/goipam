package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// SearchResult holds combined search results across multiple resource types.
type SearchResult struct {
	Addresses []*Address `json:"addresses"`
	Subnets   []*Subnet  `json:"subnets"`
	Devices   []*Device  `json:"devices"`
}

// SearchRepository handles cross-resource search data access.
type SearchRepository struct {
	q *sqlc.Queries
}

// NewSearchRepository creates a new SearchRepository.
func NewSearchRepository(q *sqlc.Queries) *SearchRepository {
	return &SearchRepository{q: q}
}

// SearchAddresses searches addresses by keyword.
func (r *SearchRepository) SearchAddresses(ctx context.Context, keyword string, limit, offset int32) ([]*Address, error) {
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
	items := make([]*Address, len(rows))
	for i, row := range rows {
		items[i] = sqlcAddressToAddress(row)
	}
	return items, nil
}

// SearchSubnets searches subnets by keyword.
func (r *SearchRepository) SearchSubnets(ctx context.Context, keyword string, limit int32) ([]*Subnet, error) {
	pattern := "%" + keyword + "%"
	rows, err := r.q.SearchSubnets(ctx, sqlc.SearchSubnetsParams{
		Name:        pattern,
		Cidr:        pattern,
		Description: sql.NullString{String: pattern, Valid: true},
		Limit:       limit,
	})
	if err != nil {
		return nil, fmt.Errorf("search subnets: %w", err)
	}
	items := make([]*Subnet, len(rows))
	for i, row := range rows {
		items[i] = sqlcSubnetToSubnet(row)
	}
	return items, nil
}

// SearchDevices searches devices by keyword.
func (r *SearchRepository) SearchDevices(ctx context.Context, keyword string, limit int32) ([]*Device, error) {
	pattern := "%" + keyword + "%"
	rows, err := r.q.SearchDevices(ctx, sqlc.SearchDevicesParams{
		Name:     pattern,
		Hostname: pattern,
		MgmtIp:   pattern,
		Limit:    limit,
	})
	if err != nil {
		return nil, fmt.Errorf("search devices: %w", err)
	}
	items := make([]*Device, len(rows))
	for i, row := range rows {
		items[i] = sqlcDeviceToDevice(row)
	}
	return items, nil
}

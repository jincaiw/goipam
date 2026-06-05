package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Subnet represents a subnet domain object.
type Subnet struct {
	ID               int64     `json:"id"`
	SectionID        int64     `json:"section_id"`
	ParentID         *int64    `json:"parent_id"`
	Cidr             string    `json:"cidr"`
	IpVersion        uint8     `json:"ip_version"`
	Network          []byte    `json:"network"`
	PrefixLen        uint8     `json:"prefix_len"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	VlanID           *int64    `json:"vlan_id"`
	VrfID            *int64    `json:"vrf_id"`
	LocationID       *int64    `json:"location_id"`
	NameserverID     *int64    `json:"nameserver_id"`
	ScanEnabled      bool      `json:"scan_enabled"`
	DiscoverEnabled  bool      `json:"discover_enabled"`
	ThresholdPercent int32     `json:"threshold_percent"`
	AllowRequest     bool      `json:"allow_request"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// SubnetRepository handles subnet data access.
type SubnetRepository struct {
	q  *sqlc.Queries
	db *sql.DB
}

// NewSubnetRepository creates a new SubnetRepository.
func NewSubnetRepository(q *sqlc.Queries, db *sql.DB) *SubnetRepository {
	return &SubnetRepository{q: q, db: db}
}

// GetByID retrieves a subnet by ID.
func (r *SubnetRepository) GetByID(ctx context.Context, id int64) (*Subnet, error) {
	row, err := r.q.GetSubnetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get subnet by id: %w", err)
	}
	return sqlcSubnetToSubnet(row), nil
}

// List retrieves a paginated list of subnets in a section.
func (r *SubnetRepository) List(ctx context.Context, sectionID int64, limit, offset int32) ([]*Subnet, error) {
	rows, err := r.q.ListSubnets(ctx, sqlc.ListSubnetsParams{
		SectionID: sectionID,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list subnets: %w", err)
	}
	subnets := make([]*Subnet, len(rows))
	for i, row := range rows {
		subnets[i] = sqlcSubnetToSubnet(row)
	}
	return subnets, nil
}

// Count returns the total number of subnets in a section.
func (r *SubnetRepository) Count(ctx context.Context, sectionID int64) (int64, error) {
	return r.q.CountSubnets(ctx, sectionID)
}

// Create creates a new subnet and returns the ID.
func (r *SubnetRepository) Create(ctx context.Context, sectionID int64, parentID *int64, cidr string, ipVersion uint8, network []byte, prefixLen uint8, name, description string, vlanID, vrfID, locationID, nameserverID *int64, scanEnabled, discoverEnabled bool, thresholdPercent int32, allowRequest bool, status string) (int64, error) {
	result, err := r.q.CreateSubnet(ctx, sqlc.CreateSubnetParams{
		SectionID:        sectionID,
		ParentID:         NullInt64(parentID),
		Cidr:             cidr,
		IpVersion:        ipVersion,
		Network:          network,
		PrefixLen:        prefixLen,
		Name:             name,
		Description:      NullString(&description),
		VlanID:           NullInt64(vlanID),
		VrfID:            NullInt64(vrfID),
		LocationID:       NullInt64(locationID),
		NameserverID:     NullInt64(nameserverID),
		ScanEnabled:      scanEnabled,
		DiscoverEnabled:  discoverEnabled,
		ThresholdPercent: thresholdPercent,
		AllowRequest:     allowRequest,
		Status:           status,
	})
	if err != nil {
		return 0, fmt.Errorf("create subnet: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a subnet.
func (r *SubnetRepository) Update(ctx context.Context, id int64, name, description string, vlanID, vrfID, locationID, nameserverID *int64, scanEnabled, discoverEnabled bool, thresholdPercent int32, allowRequest bool, status string) error {
	return r.q.UpdateSubnet(ctx, sqlc.UpdateSubnetParams{
		Name:             name,
		Description:      NullString(&description),
		VlanID:           NullInt64(vlanID),
		VrfID:            NullInt64(vrfID),
		LocationID:       NullInt64(locationID),
		NameserverID:     NullInt64(nameserverID),
		ScanEnabled:      scanEnabled,
		DiscoverEnabled:  discoverEnabled,
		ThresholdPercent: thresholdPercent,
		AllowRequest:     allowRequest,
		Status:           status,
		ID:               id,
	})
}

// Delete deletes a subnet.
func (r *SubnetRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteSubnet(ctx, id)
}

// UpdateCIDR updates a subnet's CIDR-related fields (cidr, ip_version, network, prefix_len).
func (r *SubnetRepository) UpdateCIDR(ctx context.Context, id int64, cidr string, ipVersion uint8, network []byte, prefixLen uint8) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE subnets SET cidr = ?, ip_version = ?, network = ?, prefix_len = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?",
		cidr, ipVersion, network, prefixLen, id,
	)
	if err != nil {
		return fmt.Errorf("update subnet cidr: %w", err)
	}
	return nil
}

// GetChildSubnets retrieves child subnets by parent ID.
func (r *SubnetRepository) GetChildSubnets(ctx context.Context, parentID int64) ([]*Subnet, error) {
	rows, err := r.q.GetChildSubnets(ctx, sql.NullInt64{Int64: parentID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("get child subnets: %w", err)
	}
	subnets := make([]*Subnet, len(rows))
	for i, row := range rows {
		subnets[i] = sqlcSubnetToSubnet(row)
	}
	return subnets, nil
}

// FindOverlappingSubnets finds subnets in the same section/VRF that may overlap.
// Actual overlap detection using CIDR comparison should be done by the caller.
func (r *SubnetRepository) FindOverlappingSubnets(ctx context.Context, sectionID int64, vrfID *int64, excludeID int64) ([]*Subnet, error) {
	rows, err := r.q.FindOverlappingSubnets(ctx, sqlc.FindOverlappingSubnetsParams{
		SectionID: sectionID,
		VrfID:     NullInt64(vrfID),
		ID:        excludeID,
	})
	if err != nil {
		return nil, fmt.Errorf("find overlapping subnets: %w", err)
	}
	subnets := make([]*Subnet, len(rows))
	for i, row := range rows {
		subnets[i] = sqlcSubnetToSubnet(row)
	}
	return subnets, nil
}

func sqlcSubnetToSubnet(row sqlc.Subnet) *Subnet {
	return &Subnet{
		ID:               row.ID,
		SectionID:        row.SectionID,
		ParentID:         nullInt64Ptr(row.ParentID),
		Cidr:             row.Cidr,
		IpVersion:        row.IpVersion,
		Network:          row.Network,
		PrefixLen:        row.PrefixLen,
		Name:             row.Name,
		Description:      row.Description.String,
		VlanID:           nullInt64Ptr(row.VlanID),
		VrfID:            nullInt64Ptr(row.VrfID),
		LocationID:       nullInt64Ptr(row.LocationID),
		NameserverID:     nullInt64Ptr(row.NameserverID),
		ScanEnabled:      row.ScanEnabled,
		DiscoverEnabled:  row.DiscoverEnabled,
		ThresholdPercent: row.ThresholdPercent,
		AllowRequest:     row.AllowRequest,
		Status:           row.Status,
		CreatedAt:        row.CreatedAt,
		UpdatedAt:        row.UpdatedAt,
	}
}

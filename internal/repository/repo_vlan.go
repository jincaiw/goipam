package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Vlan represents a VLAN domain object.
type Vlan struct {
	ID          int64     `json:"id"`
	L2DomainID  *int64    `json:"l2_domain_id"`
	VlanID      int32     `json:"vlan_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// L2Domain represents an L2 domain domain object.
type L2Domain struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// VlanRepository handles VLAN and L2 domain data access.
type VlanRepository struct {
	q *sqlc.Queries
}

// NewVlanRepository creates a new VlanRepository.
func NewVlanRepository(q *sqlc.Queries) *VlanRepository {
	return &VlanRepository{q: q}
}

// GetVlanByID retrieves a VLAN by ID.
func (r *VlanRepository) GetVlanByID(ctx context.Context, id int64) (*Vlan, error) {
	row, err := r.q.GetVlanByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get vlan by id: %w", err)
	}
	return sqlcVlanToVlan(row), nil
}

// ListVlans retrieves a paginated list of VLANs.
func (r *VlanRepository) ListVlans(ctx context.Context, limit, offset int32) ([]*Vlan, error) {
	rows, err := r.q.ListVlans(ctx, sqlc.ListVlansParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list vlans: %w", err)
	}
	vlans := make([]*Vlan, len(rows))
	for i, row := range rows {
		vlans[i] = sqlcVlanToVlan(row)
	}
	return vlans, nil
}

// CountVlans returns the total number of VLANs.
func (r *VlanRepository) CountVlans(ctx context.Context) (int64, error) {
	return r.q.CountVlans(ctx)
}

// CreateVlan creates a new VLAN and returns the ID.
func (r *VlanRepository) CreateVlan(ctx context.Context, l2DomainID *int64, vlanID int32, name, description string) (int64, error) {
	result, err := r.q.CreateVlan(ctx, sqlc.CreateVlanParams{
		L2DomainID:  NullInt64(l2DomainID),
		VlanID:      vlanID,
		Name:        name,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create vlan: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateVlan updates a VLAN.
func (r *VlanRepository) UpdateVlan(ctx context.Context, id int64, l2DomainID *int64, vlanID int32, name, description string) error {
	return r.q.UpdateVlan(ctx, sqlc.UpdateVlanParams{
		L2DomainID:  NullInt64(l2DomainID),
		VlanID:      vlanID,
		Name:        name,
		Description: NullString(&description),
		ID:          id,
	})
}

// DeleteVlan deletes a VLAN.
func (r *VlanRepository) DeleteVlan(ctx context.Context, id int64) error {
	return r.q.DeleteVlan(ctx, id)
}

// GetL2DomainByID retrieves an L2 domain by ID.
func (r *VlanRepository) GetL2DomainByID(ctx context.Context, id int64) (*L2Domain, error) {
	row, err := r.q.GetL2DomainByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get l2 domain by id: %w", err)
	}
	return sqlcL2DomainToL2Domain(row), nil
}

// ListL2Domains retrieves a paginated list of L2 domains.
func (r *VlanRepository) ListL2Domains(ctx context.Context, limit, offset int32) ([]*L2Domain, error) {
	rows, err := r.q.ListL2Domains(ctx, sqlc.ListL2DomainsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list l2 domains: %w", err)
	}
	domains := make([]*L2Domain, len(rows))
	for i, row := range rows {
		domains[i] = sqlcL2DomainToL2Domain(row)
	}
	return domains, nil
}

// CountL2Domains returns the total number of L2 domains.
func (r *VlanRepository) CountL2Domains(ctx context.Context) (int64, error) {
	return r.q.CountL2Domains(ctx)
}

// CreateL2Domain creates a new L2 domain and returns the ID.
func (r *VlanRepository) CreateL2Domain(ctx context.Context, name, description string) (int64, error) {
	result, err := r.q.CreateL2Domain(ctx, sqlc.CreateL2DomainParams{
		Name:        name,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create l2 domain: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateL2Domain updates an L2 domain.
func (r *VlanRepository) UpdateL2Domain(ctx context.Context, id int64, name, description string) error {
	return r.q.UpdateL2Domain(ctx, sqlc.UpdateL2DomainParams{
		Name:        name,
		Description: NullString(&description),
		ID:          id,
	})
}

// DeleteL2Domain deletes an L2 domain.
func (r *VlanRepository) DeleteL2Domain(ctx context.Context, id int64) error {
	return r.q.DeleteL2Domain(ctx, id)
}

func sqlcVlanToVlan(row sqlc.Vlan) *Vlan {
	return &Vlan{
		ID:          row.ID,
		L2DomainID:  nullInt64Ptr(row.L2DomainID),
		VlanID:      row.VlanID,
		Name:        row.Name,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func sqlcL2DomainToL2Domain(row sqlc.L2Domain) *L2Domain {
	return &L2Domain{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

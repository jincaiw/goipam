package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Vrf represents a VRF domain object.
type Vrf struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Rd          string    `json:"rd"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// VrfRepository handles VRF data access.
type VrfRepository struct {
	q *sqlc.Queries
}

// NewVrfRepository creates a new VrfRepository.
func NewVrfRepository(q *sqlc.Queries) *VrfRepository {
	return &VrfRepository{q: q}
}

// GetByID retrieves a VRF by ID.
func (r *VrfRepository) GetByID(ctx context.Context, id int64) (*Vrf, error) {
	row, err := r.q.GetVrfByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get vrf by id: %w", err)
	}
	return sqlcVrfToVrf(row), nil
}

// List retrieves a paginated list of VRFs.
func (r *VrfRepository) List(ctx context.Context, limit, offset int32) ([]*Vrf, error) {
	rows, err := r.q.ListVrfs(ctx, sqlc.ListVrfsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list vrfs: %w", err)
	}
	vrfs := make([]*Vrf, len(rows))
	for i, row := range rows {
		vrfs[i] = sqlcVrfToVrf(row)
	}
	return vrfs, nil
}

// Count returns the total number of VRFs.
func (r *VrfRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountVrfs(ctx)
}

// Create creates a new VRF and returns the ID.
func (r *VrfRepository) Create(ctx context.Context, name, rd, description string) (int64, error) {
	result, err := r.q.CreateVrf(ctx, sqlc.CreateVrfParams{
		Name:        name,
		Rd:          rd,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create vrf: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a VRF.
func (r *VrfRepository) Update(ctx context.Context, id int64, name, rd, description string) error {
	return r.q.UpdateVrf(ctx, sqlc.UpdateVrfParams{
		Name:        name,
		Rd:          rd,
		Description: NullString(&description),
		ID:          id,
	})
}

// Delete deletes a VRF.
func (r *VrfRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteVrf(ctx, id)
}

func sqlcVrfToVrf(row sqlc.Vrf) *Vrf {
	return &Vrf{
		ID:          row.ID,
		Name:        row.Name,
		Rd:          row.Rd,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

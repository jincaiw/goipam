package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Section represents a section domain object.
type Section struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    *int64    `json:"parent_id"`
	StrictMode  bool      `json:"strict_mode"`
	ShowVlan    bool      `json:"show_vlan"`
	ShowVrf     bool      `json:"show_vrf"`
	SubnetOrder string    `json:"subnet_order"`
	SortOrder   int32     `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SectionRepository handles section data access.
type SectionRepository struct {
	q *sqlc.Queries
}

// NewSectionRepository creates a new SectionRepository.
func NewSectionRepository(q *sqlc.Queries) *SectionRepository {
	return &SectionRepository{q: q}
}

// GetByID retrieves a section by ID.
func (r *SectionRepository) GetByID(ctx context.Context, id int64) (*Section, error) {
	row, err := r.q.GetSectionByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get section by id: %w", err)
	}
	return sqlcSectionToSection(row), nil
}

// List retrieves a paginated list of sections.
func (r *SectionRepository) List(ctx context.Context, limit, offset int32) ([]*Section, error) {
	rows, err := r.q.ListSections(ctx, sqlc.ListSectionsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list sections: %w", err)
	}
	sections := make([]*Section, len(rows))
	for i, row := range rows {
		sections[i] = sqlcSectionToSection(row)
	}
	return sections, nil
}

// Count returns the total number of sections.
func (r *SectionRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountSections(ctx)
}

// Create creates a new section and returns the ID.
func (r *SectionRepository) Create(ctx context.Context, name, description string, parentID *int64, strictMode, showVlan, showVrf bool, subnetOrder string, sortOrder int32) (int64, error) {
	result, err := r.q.CreateSection(ctx, sqlc.CreateSectionParams{
		Name:        name,
		Description: NullString(&description),
		ParentID:    NullInt64(parentID),
		StrictMode:  strictMode,
		ShowVlan:    showVlan,
		ShowVrf:     showVrf,
		SubnetOrder: subnetOrder,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return 0, fmt.Errorf("create section: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a section.
func (r *SectionRepository) Update(ctx context.Context, id int64, name, description string, parentID *int64, strictMode, showVlan, showVrf bool, subnetOrder string, sortOrder int32) error {
	return r.q.UpdateSection(ctx, sqlc.UpdateSectionParams{
		Name:        name,
		Description: NullString(&description),
		ParentID:    NullInt64(parentID),
		StrictMode:  strictMode,
		ShowVlan:    showVlan,
		ShowVrf:     showVrf,
		SubnetOrder: subnetOrder,
		SortOrder:   sortOrder,
		ID:          id,
	})
}

// Delete deletes a section.
func (r *SectionRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteSection(ctx, id)
}

// GetSubSections retrieves child sections by parent ID.
func (r *SectionRepository) GetSubSections(ctx context.Context, parentID int64) ([]*Section, error) {
	rows, err := r.q.GetSubSections(ctx, sql.NullInt64{Int64: parentID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("get sub sections: %w", err)
	}
	sections := make([]*Section, len(rows))
	for i, row := range rows {
		sections[i] = sqlcSectionToSection(row)
	}
	return sections, nil
}

func sqlcSectionToSection(row sqlc.Section) *Section {
	return &Section{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description.String,
		ParentID:    nullInt64Ptr(row.ParentID),
		StrictMode:  row.StrictMode,
		ShowVlan:    row.ShowVlan,
		ShowVrf:     row.ShowVrf,
		SubnetOrder: row.SubnetOrder,
		SortOrder:   row.SortOrder,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func nullInt64Ptr(v sql.NullInt64) *int64 {
	if v.Valid {
		return &v.Int64
	}
	return nil
}

func nullStringPtr(v sql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

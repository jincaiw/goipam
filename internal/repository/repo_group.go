package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Group represents a group domain object.
type Group struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// GroupRepository handles group data access.
type GroupRepository struct {
	q *sqlc.Queries
}

// NewGroupRepository creates a new GroupRepository.
func NewGroupRepository(q *sqlc.Queries) *GroupRepository {
	return &GroupRepository{q: q}
}

// GetByID retrieves a group by ID.
func (r *GroupRepository) GetByID(ctx context.Context, id int64) (*Group, error) {
	row, err := r.q.GetGroupByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get group by id: %w", err)
	}
	return sqlcGroupToGroup(row), nil
}

// List retrieves all groups.
func (r *GroupRepository) List(ctx context.Context) ([]*Group, error) {
	rows, err := r.q.ListGroups(ctx)
	if err != nil {
		return nil, fmt.Errorf("list groups: %w", err)
	}
	groups := make([]*Group, len(rows))
	for i, row := range rows {
		groups[i] = sqlcGroupToGroup(row)
	}
	return groups, nil
}

// Create creates a new group and returns the ID.
func (r *GroupRepository) Create(ctx context.Context, name, description string) (int64, error) {
	result, err := r.q.CreateGroup(ctx, sqlc.CreateGroupParams{
		Name:        name,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create group: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a group.
func (r *GroupRepository) Update(ctx context.Context, id int64, name, description string) error {
	return r.q.UpdateGroup(ctx, sqlc.UpdateGroupParams{
		Name:        name,
		Description: NullString(&description),
		ID:          id,
	})
}

// Delete deletes a group.
func (r *GroupRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteGroup(ctx, id)
}

func sqlcGroupToGroup(row sqlc.Group) *Group {
	return &Group{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

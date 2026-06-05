package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Role represents a role domain object.
type Role struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RoleRepository handles role data access.
type RoleRepository struct {
	q *sqlc.Queries
}

// NewRoleRepository creates a new RoleRepository.
func NewRoleRepository(q *sqlc.Queries) *RoleRepository {
	return &RoleRepository{q: q}
}

// GetByID retrieves a role by ID.
func (r *RoleRepository) GetByID(ctx context.Context, id int64) (*Role, error) {
	row, err := r.q.GetRoleByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get role by id: %w", err)
	}
	return sqlcRoleToRole(row), nil
}

// List retrieves all roles.
func (r *RoleRepository) List(ctx context.Context) ([]*Role, error) {
	rows, err := r.q.ListRoles(ctx)
	if err != nil {
		return nil, fmt.Errorf("list roles: %w", err)
	}
	roles := make([]*Role, len(rows))
	for i, row := range rows {
		roles[i] = sqlcRoleToRole(row)
	}
	return roles, nil
}

// Create creates a new role and returns the ID.
func (r *RoleRepository) Create(ctx context.Context, name, description string) (int64, error) {
	result, err := r.q.CreateRole(ctx, sqlc.CreateRoleParams{
		Name:        name,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create role: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a role.
func (r *RoleRepository) Update(ctx context.Context, id int64, name, description string) error {
	return r.q.UpdateRole(ctx, sqlc.UpdateRoleParams{
		Name:        name,
		Description: NullString(&description),
		ID:          id,
	})
}

// Delete deletes a role.
func (r *RoleRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteRole(ctx, id)
}

func sqlcRoleToRole(row sqlc.Role) *Role {
	return &Role{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

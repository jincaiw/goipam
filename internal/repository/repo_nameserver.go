package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Nameserver represents a nameserver domain object.
type Nameserver struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Server      string    `json:"server"`
	Port        int32     `json:"port"`
	TimeoutMs   int32     `json:"timeout_ms"`
	Enabled     bool      `json:"enabled"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NameserverRepository handles nameserver data access.
type NameserverRepository struct {
	q *sqlc.Queries
}

// NewNameserverRepository creates a new NameserverRepository.
func NewNameserverRepository(q *sqlc.Queries) *NameserverRepository {
	return &NameserverRepository{q: q}
}

// GetByID retrieves a nameserver by ID.
func (r *NameserverRepository) GetByID(ctx context.Context, id int64) (*Nameserver, error) {
	row, err := r.q.GetNameserverByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get nameserver by id: %w", err)
	}
	return sqlcNameserverToNameserver(row), nil
}

// List retrieves a paginated list of nameservers.
func (r *NameserverRepository) List(ctx context.Context, limit, offset int32) ([]*Nameserver, error) {
	rows, err := r.q.ListNameservers(ctx, sqlc.ListNameserversParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list nameservers: %w", err)
	}
	items := make([]*Nameserver, len(rows))
	for i, row := range rows {
		items[i] = sqlcNameserverToNameserver(row)
	}
	return items, nil
}

// Count returns the total number of nameservers.
func (r *NameserverRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountNameservers(ctx)
}

// Create creates a new nameserver and returns the ID.
func (r *NameserverRepository) Create(ctx context.Context, name, nsType, server string, port, timeoutMs int32, enabled bool, description string) (int64, error) {
	result, err := r.q.CreateNameserver(ctx, sqlc.CreateNameserverParams{
		Name:        name,
		Type:        nsType,
		Server:      server,
		Port:        port,
		TimeoutMs:   timeoutMs,
		Enabled:     enabled,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create nameserver: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a nameserver.
func (r *NameserverRepository) Update(ctx context.Context, id int64, name, nsType, server string, port, timeoutMs int32, enabled bool, description string) error {
	return r.q.UpdateNameserver(ctx, sqlc.UpdateNameserverParams{
		Name:        name,
		Type:        nsType,
		Server:      server,
		Port:        port,
		TimeoutMs:   timeoutMs,
		Enabled:     enabled,
		Description: NullString(&description),
		ID:          id,
	})
}

// Delete deletes a nameserver.
func (r *NameserverRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteNameserver(ctx, id)
}

func sqlcNameserverToNameserver(row sqlc.Nameserver) *Nameserver {
	return &Nameserver{
		ID:          row.ID,
		Name:        row.Name,
		Type:        row.Type,
		Server:      row.Server,
		Port:        row.Port,
		TimeoutMs:   row.TimeoutMs,
		Enabled:     row.Enabled,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

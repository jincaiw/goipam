package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Tag represents a tag domain object.
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TagRepository handles tag data access.
type TagRepository struct {
	q *sqlc.Queries
}

// NewTagRepository creates a new TagRepository.
func NewTagRepository(q *sqlc.Queries) *TagRepository {
	return &TagRepository{q: q}
}

// GetByID retrieves a tag by ID.
func (r *TagRepository) GetByID(ctx context.Context, id int64) (*Tag, error) {
	row, err := r.q.GetTagByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get tag by id: %w", err)
	}
	return sqlcTagToTag(row), nil
}

// List retrieves all tags.
func (r *TagRepository) List(ctx context.Context) ([]*Tag, error) {
	rows, err := r.q.ListTags(ctx)
	if err != nil {
		return nil, fmt.Errorf("list tags: %w", err)
	}
	tags := make([]*Tag, len(rows))
	for i, row := range rows {
		tags[i] = sqlcTagToTag(row)
	}
	return tags, nil
}

// Create creates a new tag and returns the ID.
func (r *TagRepository) Create(ctx context.Context, name, color string) (int64, error) {
	result, err := r.q.CreateTag(ctx, sqlc.CreateTagParams{
		Name:  name,
		Color: color,
	})
	if err != nil {
		return 0, fmt.Errorf("create tag: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a tag.
func (r *TagRepository) Update(ctx context.Context, id int64, name, color string) error {
	return r.q.UpdateTag(ctx, sqlc.UpdateTagParams{
		Name:  name,
		Color: color,
		ID:    id,
	})
}

// Delete deletes a tag.
func (r *TagRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteTag(ctx, id)
}

func sqlcTagToTag(row sqlc.Tag) *Tag {
	return &Tag{
		ID:        row.ID,
		Name:      row.Name,
		Color:     row.Color,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

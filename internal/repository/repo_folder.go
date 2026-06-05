package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Folder represents a folder domain object.
type Folder struct {
	ID          int64     `json:"id"`
	SectionID   int64     `json:"section_id"`
	ParentID    *int64    `json:"parent_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	SortOrder   int32     `json:"sort_order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// FolderRepository handles folder data access.
type FolderRepository struct {
	q *sqlc.Queries
}

// NewFolderRepository creates a new FolderRepository.
func NewFolderRepository(q *sqlc.Queries) *FolderRepository {
	return &FolderRepository{q: q}
}

// GetByID retrieves a folder by ID.
func (r *FolderRepository) GetByID(ctx context.Context, id int64) (*Folder, error) {
	row, err := r.q.GetFolderByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get folder by id: %w", err)
	}
	return sqlcFolderToFolder(row), nil
}

// List retrieves a paginated list of folders in a section.
func (r *FolderRepository) List(ctx context.Context, sectionID int64, limit, offset int32) ([]*Folder, error) {
	rows, err := r.q.ListFolders(ctx, sqlc.ListFoldersParams{
		SectionID: sectionID,
		Limit:     limit,
		Offset:    offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list folders: %w", err)
	}
	folders := make([]*Folder, len(rows))
	for i, row := range rows {
		folders[i] = sqlcFolderToFolder(row)
	}
	return folders, nil
}

// Count returns the total number of folders in a section.
func (r *FolderRepository) Count(ctx context.Context, sectionID int64) (int64, error) {
	return r.q.CountFolders(ctx, sectionID)
}

// Create creates a new folder and returns the ID.
func (r *FolderRepository) Create(ctx context.Context, sectionID int64, parentID *int64, name, description string, sortOrder int32) (int64, error) {
	result, err := r.q.CreateFolder(ctx, sqlc.CreateFolderParams{
		SectionID:   sectionID,
		ParentID:    NullInt64(parentID),
		Name:        name,
		Description: NullString(&description),
		SortOrder:   sortOrder,
	})
	if err != nil {
		return 0, fmt.Errorf("create folder: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a folder.
func (r *FolderRepository) Update(ctx context.Context, id int64, name, description string, sortOrder int32) error {
	return r.q.UpdateFolder(ctx, sqlc.UpdateFolderParams{
		Name:        name,
		Description: NullString(&description),
		SortOrder:   sortOrder,
		ID:          id,
	})
}

// Delete deletes a folder.
func (r *FolderRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteFolder(ctx, id)
}

// GetSubFolders retrieves child folders by parent ID.
func (r *FolderRepository) GetSubFolders(ctx context.Context, parentID int64) ([]*Folder, error) {
	rows, err := r.q.GetSubFolders(ctx, sql.NullInt64{Int64: parentID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("get sub folders: %w", err)
	}
	folders := make([]*Folder, len(rows))
	for i, row := range rows {
		folders[i] = sqlcFolderToFolder(row)
	}
	return folders, nil
}

func sqlcFolderToFolder(row sqlc.Folder) *Folder {
	return &Folder{
		ID:          row.ID,
		SectionID:   row.SectionID,
		ParentID:    nullInt64Ptr(row.ParentID),
		Name:        row.Name,
		Description: row.Description.String,
		SortOrder:   row.SortOrder,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

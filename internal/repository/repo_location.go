package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// Location represents a location domain object.
type Location struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Room        string    `json:"room"`
	RowName     string    `json:"row_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Rack represents a rack domain object.
type Rack struct {
	ID          int64     `json:"id"`
	LocationID  int64     `json:"location_id"`
	Name        string    `json:"name"`
	SizeU       int32     `json:"size_u"`
	RowName     string    `json:"row_name"`
	Position    string    `json:"position"`
	Direction   string    `json:"direction"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// LocationRepository handles location and rack data access.
type LocationRepository struct {
	q *sqlc.Queries
}

// NewLocationRepository creates a new LocationRepository.
func NewLocationRepository(q *sqlc.Queries) *LocationRepository {
	return &LocationRepository{q: q}
}

// GetLocationByID retrieves a location by ID.
func (r *LocationRepository) GetLocationByID(ctx context.Context, id int64) (*Location, error) {
	row, err := r.q.GetLocationByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get location by id: %w", err)
	}
	return sqlcLocationToLocation(row), nil
}

// ListLocations retrieves a paginated list of locations.
func (r *LocationRepository) ListLocations(ctx context.Context, limit, offset int32) ([]*Location, error) {
	rows, err := r.q.ListLocations(ctx, sqlc.ListLocationsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list locations: %w", err)
	}
	locations := make([]*Location, len(rows))
	for i, row := range rows {
		locations[i] = sqlcLocationToLocation(row)
	}
	return locations, nil
}

// CountLocations returns the total number of locations.
func (r *LocationRepository) CountLocations(ctx context.Context) (int64, error) {
	return r.q.CountLocations(ctx)
}

// CreateLocation creates a new location and returns the ID.
func (r *LocationRepository) CreateLocation(ctx context.Context, name, address, room, rowName, description string) (int64, error) {
	result, err := r.q.CreateLocation(ctx, sqlc.CreateLocationParams{
		Name:        name,
		Address:     address,
		Room:        room,
		RowName:     rowName,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create location: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateLocation updates a location.
func (r *LocationRepository) UpdateLocation(ctx context.Context, id int64, name, address, room, rowName, description string) error {
	return r.q.UpdateLocation(ctx, sqlc.UpdateLocationParams{
		Name:        name,
		Address:     address,
		Room:        room,
		RowName:     rowName,
		Description: NullString(&description),
		ID:          id,
	})
}

// DeleteLocation deletes a location.
func (r *LocationRepository) DeleteLocation(ctx context.Context, id int64) error {
	return r.q.DeleteLocation(ctx, id)
}

// GetRackByID retrieves a rack by ID.
func (r *LocationRepository) GetRackByID(ctx context.Context, id int64) (*Rack, error) {
	row, err := r.q.GetRackByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get rack by id: %w", err)
	}
	return sqlcRackToRack(row), nil
}

// ListRacks retrieves a paginated list of racks. If locationID is 0, all racks are returned.
func (r *LocationRepository) ListRacks(ctx context.Context, locationID int64, limit, offset int32) ([]*Rack, error) {
	if locationID == 0 {
		rows, err := r.q.ListAllRacks(ctx, sqlc.ListAllRacksParams{
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("list all racks: %w", err)
		}
		racks := make([]*Rack, len(rows))
		for i, row := range rows {
			racks[i] = sqlcRackToRack(row)
		}
		return racks, nil
	}

	rows, err := r.q.ListRacks(ctx, sqlc.ListRacksParams{
		LocationID: locationID,
		Limit:      limit,
		Offset:     offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list racks: %w", err)
	}
	racks := make([]*Rack, len(rows))
	for i, row := range rows {
		racks[i] = sqlcRackToRack(row)
	}
	return racks, nil
}

// CountRacks returns the total number of racks. If locationID is 0, all racks are counted.
func (r *LocationRepository) CountRacks(ctx context.Context, locationID int64) (int64, error) {
	if locationID == 0 {
		return r.q.CountAllRacks(ctx)
	}
	return r.q.CountRacks(ctx, locationID)
}

// CreateRack creates a new rack and returns the ID.
func (r *LocationRepository) CreateRack(ctx context.Context, locationID int64, name string, sizeU int32, rowName, position, direction, description string) (int64, error) {
	result, err := r.q.CreateRack(ctx, sqlc.CreateRackParams{
		LocationID:  locationID,
		Name:        name,
		SizeU:       sizeU,
		RowName:     rowName,
		Position:    position,
		Direction:   direction,
		Description: NullString(&description),
	})
	if err != nil {
		return 0, fmt.Errorf("create rack: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// UpdateRack updates a rack.
func (r *LocationRepository) UpdateRack(ctx context.Context, id int64, name string, sizeU int32, rowName, position, direction, description string) error {
	return r.q.UpdateRack(ctx, sqlc.UpdateRackParams{
		Name:        name,
		SizeU:       sizeU,
		RowName:     rowName,
		Position:    position,
		Direction:   direction,
		Description: NullString(&description),
		ID:          id,
	})
}

// DeleteRack deletes a rack.
func (r *LocationRepository) DeleteRack(ctx context.Context, id int64) error {
	return r.q.DeleteRack(ctx, id)
}

func sqlcLocationToLocation(row sqlc.Location) *Location {
	return &Location{
		ID:          row.ID,
		Name:        row.Name,
		Address:     row.Address,
		Room:        row.Room,
		RowName:     row.RowName,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func sqlcRackToRack(row sqlc.Rack) *Rack {
	return &Rack{
		ID:          row.ID,
		LocationID:  row.LocationID,
		Name:        row.Name,
		SizeU:       row.SizeU,
		RowName:     row.RowName,
		Position:    row.Position,
		Direction:   row.Direction,
		Description: row.Description.String,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

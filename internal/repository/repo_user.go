package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// User represents a user domain object (without password hash).
type User struct {
	ID               int64      `json:"id"`
	Username         string     `json:"username"`
	DisplayName      string     `json:"display_name"`
	Email            string     `json:"email"`
	Status           string     `json:"status"`
	FailedLoginCount int32      `json:"failed_login_count"`
	LockedUntil      *time.Time `json:"locked_until"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
}

// UserWithPassword includes the password hash for authentication.
type UserWithPassword struct {
	User
	PasswordHash string `json:"-"`
}

// UserRepository handles user data access.
type UserRepository struct {
	q *sqlc.Queries
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(q *sqlc.Queries) *UserRepository {
	return &UserRepository{q: q}
}

// GetByID retrieves a user by ID.
func (r *UserRepository) GetByID(ctx context.Context, id int64) (*User, error) {
	row, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}
	return sqlcUserToUser(row), nil
}

// GetByUsername retrieves a user by username (with password hash).
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*UserWithPassword, error) {
	row, err := r.q.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}
	u := sqlcUserToUser(row)
	return &UserWithPassword{
		User:         *u,
		PasswordHash: row.PasswordHash,
	}, nil
}

// List retrieves a paginated list of users.
func (r *UserRepository) List(ctx context.Context, limit, offset int32) ([]*User, error) {
	rows, err := r.q.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}
	users := make([]*User, len(rows))
	for i, row := range rows {
		users[i] = &User{
			ID:          row.ID,
			Username:    row.Username,
			DisplayName: row.DisplayName,
			Email:       row.Email,
			Status:      row.Status,
			LastLoginAt: nullTimePtr(row.LastLoginAt),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}
	}
	return users, nil
}

// Count returns the total number of users.
func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	return r.q.CountUsers(ctx)
}

// Create creates a new user and returns the ID.
func (r *UserRepository) Create(ctx context.Context, username, passwordHash, displayName, email, status string) (int64, error) {
	result, err := r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Username:     username,
		PasswordHash: passwordHash,
		DisplayName:  displayName,
		Email:        email,
		Status:       status,
	})
	if err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get last insert id: %w", err)
	}
	return id, nil
}

// Update updates a user.
func (r *UserRepository) Update(ctx context.Context, id int64, displayName, email, status string) error {
	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		DisplayName: displayName,
		Email:       email,
		Status:      status,
		ID:          id,
	})
}

// UpdatePassword updates a user's password hash.
func (r *UserRepository) UpdatePassword(ctx context.Context, id int64, passwordHash string) error {
	return r.q.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		PasswordHash: passwordHash,
		ID:           id,
	})
}

// UpdateLoginInfo updates login-related fields.
func (r *UserRepository) UpdateLoginInfo(ctx context.Context, id int64, failedCount int32, lockedUntil sql.NullTime) error {
	return r.q.UpdateUserLoginInfo(ctx, sqlc.UpdateUserLoginInfoParams{
		FailedLoginCount: failedCount,
		LockedUntil:      lockedUntil,
		ID:               id,
	})
}

// Delete deletes a user.
func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.q.DeleteUser(ctx, id)
}

func sqlcUserToUser(row sqlc.User) *User {
	return &User{
		ID:               row.ID,
		Username:         row.Username,
		DisplayName:      row.DisplayName,
		Email:            row.Email,
		Status:           row.Status,
		FailedLoginCount: row.FailedLoginCount,
		LockedUntil:      nullTimePtr(row.LockedUntil),
		LastLoginAt:      nullTimePtr(row.LastLoginAt),
		CreatedAt:        row.CreatedAt,
		UpdatedAt:        row.UpdatedAt,
	}
}

func nullTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

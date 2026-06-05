package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// SystemSetting represents a system setting.
type SystemSetting struct {
	ID        int64     `json:"id"`
	Category  string    `json:"category"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SettingsRepository handles system settings data access.
type SettingsRepository struct {
	db *sql.DB
}

// NewSettingsRepository creates a new SettingsRepository.
func NewSettingsRepository(db *sql.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

// Get retrieves a setting by category and key.
func (r *SettingsRepository) Get(ctx context.Context, category, key string) (string, error) {
	var value string
	err := r.db.QueryRowContext(ctx,
		`SELECT value FROM system_settings WHERE category = ? AND "key" = ?`,
		category, key,
	).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", fmt.Errorf("get setting: %w", err)
	}
	return value, nil
}

// GetByCategory retrieves all settings in a category.
func (r *SettingsRepository) GetByCategory(ctx context.Context, category string) (map[string]string, error) {
	rows, err := r.db.QueryContext(ctx,
		`SELECT "key", value FROM system_settings WHERE category = ? ORDER BY "key"`,
		category,
	)
	if err != nil {
		return nil, fmt.Errorf("get settings by category: %w", err)
	}
	defer rows.Close()

	result := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, fmt.Errorf("scan setting: %w", err)
		}
		result[key] = value
	}
	return result, rows.Err()
}

// Upsert creates or updates a setting.
func (r *SettingsRepository) Upsert(ctx context.Context, category, key, value string) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO system_settings (category, "key", value, updated_at)
		 VALUES (?, ?, ?, CURRENT_TIMESTAMP)
		 ON CONFLICT (category, "key") DO UPDATE SET value = excluded.value, updated_at = CURRENT_TIMESTAMP`,
		category, key, value,
	)
	if err != nil {
		return fmt.Errorf("upsert setting: %w", err)
	}
	return nil
}

// UpsertBatch creates or updates multiple settings in a category.
func (r *SettingsRepository) UpsertBatch(ctx context.Context, category string, settings map[string]string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	for key, value := range settings {
		_, err := tx.ExecContext(ctx,
			`INSERT INTO system_settings (category, "key", value, updated_at)
			 VALUES (?, ?, ?, CURRENT_TIMESTAMP)
			 ON CONFLICT (category, "key") DO UPDATE SET value = excluded.value, updated_at = CURRENT_TIMESTAMP`,
			category, key, value,
		)
		if err != nil {
			return fmt.Errorf("upsert setting %s.%s: %w", category, key, err)
		}
	}

	return tx.Commit()
}

// Delete removes a setting.
func (r *SettingsRepository) Delete(ctx context.Context, category, key string) error {
	_, err := r.db.ExecContext(ctx,
		`DELETE FROM system_settings WHERE category = ? AND "key" = ?`,
		category, key,
	)
	if err != nil {
		return fmt.Errorf("delete setting: %w", err)
	}
	return nil
}

package db

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"os"

	"github.com/pressly/goose/v3"

	"github.com/jincaiw/goipam/internal/config"
)

// embeddedMigrations makes release binaries independent from a runtime
// migrations/ directory. The release build script refreshes this directory.
//
//go:embed embedded_migrations/mysql/*.sql embedded_migrations/sqlite/*.sql
var embeddedMigrations embed.FS

// RunMigrations executes database migrations for the configured driver.
func RunMigrations(db *sql.DB, cfg *config.DatabaseConfig) error {
	var migrationsDir string
	switch cfg.Driver {
	case "sqlite":
		migrationsDir = "migrations/sqlite"
	case "mysql":
		migrationsDir = "migrations/mysql"
	default:
		return fmt.Errorf("unsupported driver for migrations: %s", cfg.Driver)
	}

	slog.Info("running database migrations", "driver", cfg.Driver, "dir", migrationsDir)

	if err := goose.SetDialect(cfg.Driver); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}
	migrationBase, migrationDir, err := migrationSource(migrationsDir, cfg.Driver)
	if err != nil {
		return err
	}
	goose.SetBaseFS(migrationBase)
	defer goose.SetBaseFS(nil)

	if err := goose.Up(db, migrationDir); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	version, err := goose.GetDBVersion(db)
	if err != nil {
		return fmt.Errorf("get db version: %w", err)
	}

	slog.Info("database migrations complete", "version", version)
	return nil
}

func migrationSource(filesystemDir, driver string) (fs.FS, string, error) {
	if _, err := os.Stat(filesystemDir); err == nil {
		return nil, filesystemDir, nil
	}

	embedded, err := fs.Sub(embeddedMigrations, "embedded_migrations")
	if err != nil {
		return nil, "", fmt.Errorf("load embedded migrations: %w", err)
	}
	return embedded, driver, nil
}

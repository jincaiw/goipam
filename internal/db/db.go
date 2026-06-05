package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/glebarez/sqlite"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jincaiw/goipam/internal/config"
)

// Connect opens a database connection based on the configuration.
func Connect(cfg *config.DatabaseConfig) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch cfg.Driver {
	case "sqlite":
		if err := ensureSQLiteDir(cfg.DSN); err != nil {
			return nil, err
		}
		db, err = sql.Open("sqlite", cfg.DSN)
		if err != nil {
			return nil, fmt.Errorf("open sqlite: %w", err)
		}
		// SQLite pragmas for better performance and safety
		pragmas := []string{
			"PRAGMA journal_mode=WAL",
			"PRAGMA busy_timeout=5000",
			"PRAGMA foreign_keys=ON",
		}
		for _, p := range pragmas {
			if _, err := db.Exec(p); err != nil {
				return nil, fmt.Errorf("set pragma %q: %w", p, err)
			}
		}
		slog.Info("connected to sqlite", "dsn", cfg.DSN)

	case "mysql":
		db, err = sql.Open("mysql", cfg.DSN)
		if err != nil {
			return nil, fmt.Errorf("open mysql: %w", err)
		}
		slog.Info("connected to mysql", "dsn", maskDSN(cfg.DSN))

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	return db, nil
}

func ensureSQLiteDir(dsn string) error {
	path := sqlitePathFromDSN(dsn)
	if path == "" {
		return nil
	}
	dir := filepath.Dir(path)
	if dir == "." || dir == "" {
		return nil
	}
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return fmt.Errorf("create sqlite data directory %q: %w", dir, err)
	}
	return nil
}

func sqlitePathFromDSN(dsn string) string {
	if dsn == "" || dsn == ":memory:" || strings.HasPrefix(dsn, "file::memory:") {
		return ""
	}
	if strings.HasPrefix(dsn, "file:") {
		dsn = strings.TrimPrefix(dsn, "file:")
	}
	if i := strings.IndexByte(dsn, '?'); i >= 0 {
		dsn = dsn[:i]
	}
	if dsn == "" || strings.HasPrefix(dsn, "mode=") {
		return ""
	}
	return dsn
}

// maskDSN hides the password in a MySQL DSN for logging.
func maskDSN(dsn string) string {
	// Simple masking: replace everything between : and @
	i := 0
	result := make([]byte, 0, len(dsn))
	state := 0 // 0=before colon, 1=in password, 2=after at
	for _, c := range dsn {
		switch state {
		case 0:
			if c == ':' && i > 0 {
				state = 1
				result = append(result, ':')
				result = append(result, '*')
				result = append(result, '*')
				result = append(result, '*')
				continue
			}
			result = append(result, byte(c))
		case 1:
			if c == '@' {
				state = 2
				result = append(result, '@')
			}
			// skip password chars
		case 2:
			result = append(result, byte(c))
		}
		i++
	}
	return string(result)
}

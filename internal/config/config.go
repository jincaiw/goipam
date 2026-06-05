package config

import (
	"fmt"
	"os"
	"time"
)

// Config holds all application configuration.
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Security SecurityConfig `yaml:"security"`
	Scan     ScanConfig     `yaml:"scan"`
	Log      LogConfig      `yaml:"log"`
}

type ServerConfig struct {
	Listen    string `yaml:"listen"`
	PublicURL string `yaml:"public_url"`
}

type DatabaseConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type SecurityConfig struct {
	JWTSecret       string        `yaml:"jwt_secret"`
	AccessTokenTTL  time.Duration `yaml:"access_token_ttl"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl"`
}

type ScanConfig struct {
	Enabled        bool          `yaml:"enabled"`
	MaxConcurrency int           `yaml:"max_concurrency"`
	Timeout        time.Duration `yaml:"timeout"`
	Schedule       string        `yaml:"schedule"`
}

type LogConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

// DefaultConfig returns a config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Listen:    "0.0.0.0:5050",
			PublicURL: "http://localhost:5050",
		},
		Database: DatabaseConfig{
			Driver: "sqlite",
			DSN:    "./data/goipam.db",
		},
		Security: SecurityConfig{
			JWTSecret:       "change-me-in-production",
			AccessTokenTTL:  30 * time.Minute,
			RefreshTokenTTL: 168 * time.Hour,
		},
		Scan: ScanConfig{
			Enabled:        true,
			MaxConcurrency: 128,
			Timeout:        1 * time.Second,
			Schedule:       "0 */6 * * *",
		},
		Log: LogConfig{
			Level:  "info",
			Format: "json",
		},
	}
}

// Load reads configuration from a YAML file and applies environment variable overrides.
func Load(path string) (*Config, error) {
	cfg := DefaultConfig()

	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read config file: %w", err)
		}
		if err := unmarshalYAML(data, cfg); err != nil {
			return nil, fmt.Errorf("parse config file: %w", err)
		}
	}

	// Environment variable overrides
	if v := os.Getenv("GOIPAM_LISTEN"); v != "" {
		cfg.Server.Listen = v
	}
	if v := os.Getenv("GOIPAM_PUBLIC_URL"); v != "" {
		cfg.Server.PublicURL = v
	}
	if v := os.Getenv("GOIPAM_DB_DRIVER"); v != "" {
		cfg.Database.Driver = v
	}
	if v := os.Getenv("GOIPAM_DB_DSN"); v != "" {
		cfg.Database.DSN = v
	}
	if v := os.Getenv("GOIPAM_JWT_SECRET"); v != "" {
		cfg.Security.JWTSecret = v
	}
	if v := os.Getenv("GOIPAM_LOG_LEVEL"); v != "" {
		cfg.Log.Level = v
	}

	return cfg, nil
}

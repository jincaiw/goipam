package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func unmarshalYAML(data []byte, cfg *Config) error {
	// We need a temporary struct to handle duration strings like "30m", "168h"
	var raw struct {
		Server   ServerConfig   `yaml:"server"`
		Database DatabaseConfig `yaml:"database"`
		Security struct {
			JWTSecret       string `yaml:"jwt_secret"`
			AccessTokenTTL  string `yaml:"access_token_ttl"`
			RefreshTokenTTL string `yaml:"refresh_token_ttl"`
		} `yaml:"security"`
		Scan struct {
			Enabled        *bool  `yaml:"enabled"`
			MaxConcurrency int    `yaml:"max_concurrency"`
			Timeout        string `yaml:"timeout"`
			Schedule       string `yaml:"schedule"`
		} `yaml:"scan"`
		Log LogConfig `yaml:"log"`
	}

	if err := yaml.Unmarshal(data, &raw); err != nil {
		return err
	}

	if raw.Server.Listen != "" {
		cfg.Server.Listen = raw.Server.Listen
	}
	if raw.Server.PublicURL != "" {
		cfg.Server.PublicURL = raw.Server.PublicURL
	}
	if raw.Database.Driver != "" {
		cfg.Database.Driver = raw.Database.Driver
	}
	if raw.Database.DSN != "" {
		cfg.Database.DSN = raw.Database.DSN
	}
	if raw.Log.Level != "" {
		cfg.Log.Level = raw.Log.Level
	}
	if raw.Log.Format != "" {
		cfg.Log.Format = raw.Log.Format
	}
	if raw.Security.JWTSecret != "" {
		cfg.Security.JWTSecret = raw.Security.JWTSecret
	}
	if raw.Scan.Enabled != nil {
		cfg.Scan.Enabled = *raw.Scan.Enabled
	}
	if raw.Scan.MaxConcurrency != 0 {
		cfg.Scan.MaxConcurrency = raw.Scan.MaxConcurrency
	}
	if raw.Scan.Schedule != "" {
		cfg.Scan.Schedule = raw.Scan.Schedule
	}

	var err error
	if raw.Security.AccessTokenTTL != "" {
		cfg.Security.AccessTokenTTL, err = parseDuration(raw.Security.AccessTokenTTL)
		if err != nil {
			return fmt.Errorf("parse access_token_ttl: %w", err)
		}
	}
	if raw.Security.RefreshTokenTTL != "" {
		cfg.Security.RefreshTokenTTL, err = parseDuration(raw.Security.RefreshTokenTTL)
		if err != nil {
			return fmt.Errorf("parse refresh_token_ttl: %w", err)
		}
	}
	if raw.Scan.Timeout != "" {
		cfg.Scan.Timeout, err = parseDuration(raw.Scan.Timeout)
		if err != nil {
			return fmt.Errorf("parse scan.timeout: %w", err)
		}
	}

	return nil
}

func parseDuration(s string) (time.Duration, error) {
	// Support plain integer as seconds
	if v, err := strconv.Atoi(s); err == nil {
		return time.Duration(v) * time.Second, nil
	}
	// Support Go duration strings like "30m", "1s", "168h"
	return time.ParseDuration(s)
}

// Validate checks the configuration for required values.
func (c *Config) Validate() error {
	if c.Server.Listen == "" {
		return fmt.Errorf("server.listen is required")
	}
	if c.Database.Driver != "sqlite" && c.Database.Driver != "mysql" {
		return fmt.Errorf("database.driver must be 'sqlite' or 'mysql', got %q", c.Database.Driver)
	}
	if c.Database.DSN == "" {
		return fmt.Errorf("database.dsn is required")
	}
	if c.Security.JWTSecret == "" || c.Security.JWTSecret == "change-me-in-production" {
		if isProduction() {
			return fmt.Errorf("security.jwt_secret must be changed in production")
		}
	}
	if strings.TrimSpace(c.Security.JWTSecret) == "" {
		return fmt.Errorf("security.jwt_secret cannot be empty")
	}
	return nil
}

func isProduction() bool {
	return strings.EqualFold(os.Getenv("GOIPAM_ENV"), "production")
}

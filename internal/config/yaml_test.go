package config

import (
	"testing"
	"time"
)

func TestUnmarshalYAMLKeepsDefaultsForOmittedFields(t *testing.T) {
	cfg := DefaultConfig()

	data := []byte(`
server:
  listen: "127.0.0.1:5051"
security:
  jwt_secret: "test-secret"
`)

	if err := unmarshalYAML(data, cfg); err != nil {
		t.Fatalf("unmarshalYAML() error = %v", err)
	}

	if cfg.Server.Listen != "127.0.0.1:5051" {
		t.Fatalf("Server.Listen = %q, want override", cfg.Server.Listen)
	}
	if cfg.Server.PublicURL != "http://localhost:5050" {
		t.Fatalf("Server.PublicURL = %q, want default", cfg.Server.PublicURL)
	}
	if cfg.Database.Driver != "sqlite" {
		t.Fatalf("Database.Driver = %q, want default", cfg.Database.Driver)
	}
	if cfg.Database.DSN != "./data/goipam.db" {
		t.Fatalf("Database.DSN = %q, want default", cfg.Database.DSN)
	}
	if cfg.Security.AccessTokenTTL != 30*time.Minute {
		t.Fatalf("AccessTokenTTL = %v, want default", cfg.Security.AccessTokenTTL)
	}
	if cfg.Scan.MaxConcurrency != 128 {
		t.Fatalf("Scan.MaxConcurrency = %d, want default", cfg.Scan.MaxConcurrency)
	}
	if !cfg.Scan.Enabled {
		t.Fatal("Scan.Enabled = false, want default true")
	}
	if cfg.Log.Format != "json" {
		t.Fatalf("Log.Format = %q, want default", cfg.Log.Format)
	}
}

func TestUnmarshalYAMLAllowsExplicitFalseBoolOverride(t *testing.T) {
	cfg := DefaultConfig()

	data := []byte(`
scan:
  enabled: false
`)

	if err := unmarshalYAML(data, cfg); err != nil {
		t.Fatalf("unmarshalYAML() error = %v", err)
	}

	if cfg.Scan.Enabled {
		t.Fatal("Scan.Enabled = true, want explicit false override")
	}
}

func TestUnmarshalYAMLParsesDurationOverrides(t *testing.T) {
	cfg := DefaultConfig()

	data := []byte(`
security:
  access_token_ttl: "45m"
  refresh_token_ttl: "24h"
scan:
  timeout: "250ms"
`)

	if err := unmarshalYAML(data, cfg); err != nil {
		t.Fatalf("unmarshalYAML() error = %v", err)
	}

	if cfg.Security.AccessTokenTTL != 45*time.Minute {
		t.Fatalf("AccessTokenTTL = %v, want 45m", cfg.Security.AccessTokenTTL)
	}
	if cfg.Security.RefreshTokenTTL != 24*time.Hour {
		t.Fatalf("RefreshTokenTTL = %v, want 24h", cfg.Security.RefreshTokenTTL)
	}
	if cfg.Scan.Timeout != 250*time.Millisecond {
		t.Fatalf("Scan.Timeout = %v, want 250ms", cfg.Scan.Timeout)
	}
}

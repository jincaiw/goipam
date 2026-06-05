package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/jincaiw/goipam/internal/auth"
	"github.com/jincaiw/goipam/internal/db/sqlc"
)

// SeedData creates initial admin user and default roles/permissions.
func SeedData(db *sql.DB, queries *sqlc.Queries) error {
	// Check if any users exist
	var count int
	err := db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return fmt.Errorf("check users count: %w", err)
	}
	if count > 0 {
		slog.Info("database already seeded, skipping")
		return nil
	}

	slog.Info("seeding initial data...")

	// Create admin user
	adminPassword, err := initialAdminPassword()
	if err != nil {
		return err
	}
	hash, err := auth.HashPassword(adminPassword, auth.DefaultParams)
	if err != nil {
		return fmt.Errorf("hash admin password: %w", err)
	}

	result, err := db.ExecContext(context.Background(),
		"INSERT INTO users (username, password_hash, display_name, email, status) VALUES (?, ?, ?, ?, ?)",
		"admin", hash, "Administrator", "admin@goipam.local", "active",
	)
	if err != nil {
		return fmt.Errorf("create admin user: %w", err)
	}
	adminID, _ := result.LastInsertId()
	slog.Info("created admin user", "id", adminID, "username", "admin")

	// Create default roles
	roles := []struct {
		name        string
		description string
	}{
		{"super_admin", "Super Administrator - full access to all resources"},
		{"admin", "Administrator - manage all resources"},
		{"operator", "Operator - manage IPAM resources"},
		{"viewer", "Viewer - read-only access"},
	}
	for _, role := range roles {
		_, err := db.ExecContext(context.Background(),
			"INSERT INTO roles (name, description) VALUES (?, ?)",
			role.name, role.description,
		)
		if err != nil {
			slog.Warn("failed to create role", "name", role.name, "error", err)
		}
	}

	// Assign super_admin role to admin user
	_, err = db.ExecContext(context.Background(),
		"INSERT INTO user_roles (user_id, role_id) SELECT ?, id FROM roles WHERE name = ?",
		adminID, "super_admin",
	)
	if err != nil {
		slog.Warn("failed to assign super_admin role", "error", err)
	}

	// Create default permissions
	permissions := []struct {
		code        string
		description string
	}{
		// Section permissions
		{"sections.view", "View sections"},
		{"sections.create", "Create sections"},
		{"sections.edit", "Edit sections"},
		{"sections.delete", "Delete sections"},
		// Subnet permissions
		{"subnets.view", "View subnets"},
		{"subnets.create", "Create subnets"},
		{"subnets.edit", "Edit subnets"},
		{"subnets.delete", "Delete subnets"},
		// Address permissions
		{"addresses.view", "View addresses"},
		{"addresses.create", "Create addresses"},
		{"addresses.edit", "Edit addresses"},
		{"addresses.delete", "Delete addresses"},
		// VLAN permissions
		{"vlans.view", "View VLANs"},
		{"vlans.create", "Create VLANs"},
		{"vlans.edit", "Edit VLANs"},
		{"vlans.delete", "Delete VLANs"},
		// VRF permissions
		{"vrfs.view", "View VRFs"},
		{"vrfs.create", "Create VRFs"},
		{"vrfs.edit", "Edit VRFs"},
		{"vrfs.delete", "Delete VRFs"},
		// Device permissions
		{"devices.view", "View devices"},
		{"devices.create", "Create devices"},
		{"devices.edit", "Edit devices"},
		{"devices.delete", "Delete devices"},
		// Scan permissions
		{"scans.view", "View scans"},
		{"scans.create", "Create scans"},
		// User management
		{"users.view", "View users"},
		{"users.create", "Create users"},
		{"users.edit", "Edit users"},
		{"users.delete", "Delete users"},
		// Audit
		{"audit.view", "View audit logs"},
		// Settings
		{"settings.view", "View settings"},
		{"settings.edit", "Edit settings"},
	}
	for _, perm := range permissions {
		_, err := db.ExecContext(context.Background(),
			"INSERT INTO permissions (code, description) VALUES (?, ?)",
			perm.code, perm.description,
		)
		if err != nil {
			slog.Warn("failed to create permission", "code", perm.code, "error", err)
		}
	}

	// Assign all permissions to super_admin role
	_, err = db.ExecContext(context.Background(),
		"INSERT INTO role_permissions (role_id, permission_id) SELECT r.id, p.id FROM roles r, permissions p WHERE r.name = 'super_admin'",
	)
	if err != nil {
		slog.Warn("failed to assign permissions to super_admin", "error", err)
	}

	slog.Info("database seeding complete")
	return nil
}

func initialAdminPassword() (string, error) {
	password := os.Getenv("GOIPAM_ADMIN_PASSWORD")
	if password == "" {
		if strings.EqualFold(os.Getenv("GOIPAM_ENV"), "production") {
			return "", fmt.Errorf("GOIPAM_ADMIN_PASSWORD is required for first production startup")
		}
		return "admin", nil
	}
	if len(password) < 8 {
		return "", fmt.Errorf("GOIPAM_ADMIN_PASSWORD must be at least 8 characters")
	}
	return password, nil
}

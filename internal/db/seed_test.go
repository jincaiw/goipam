package db

import "testing"

func TestInitialAdminPasswordRequiresProductionOverride(t *testing.T) {
	t.Setenv("GOIPAM_ENV", "production")
	t.Setenv("GOIPAM_ADMIN_PASSWORD", "")

	if _, err := initialAdminPassword(); err == nil {
		t.Fatal("initialAdminPassword() error = nil, want production password requirement")
	}
}

func TestInitialAdminPasswordUsesDevelopmentDefault(t *testing.T) {
	t.Setenv("GOIPAM_ENV", "")
	t.Setenv("GOIPAM_ADMIN_PASSWORD", "")

	got, err := initialAdminPassword()
	if err != nil {
		t.Fatalf("initialAdminPassword() error = %v", err)
	}
	if got != "admin" {
		t.Fatalf("initialAdminPassword() = %q, want admin", got)
	}
}

func TestInitialAdminPasswordRejectsShortOverride(t *testing.T) {
	t.Setenv("GOIPAM_ADMIN_PASSWORD", "short")

	if _, err := initialAdminPassword(); err == nil {
		t.Fatal("initialAdminPassword() error = nil, want short password rejection")
	}
}

package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "test-password-123"

	hash, err := HashPassword(password, DefaultParams)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	if hash == "" {
		t.Fatal("HashPassword returned empty hash")
	}

	// Hash should contain argon2id prefix
	if len(hash) < 20 {
		t.Fatalf("Hash seems too short: %s", hash)
	}
}

func TestHashPassword_DifferentSalts(t *testing.T) {
	password := "same-password"

	hash1, _ := HashPassword(password, DefaultParams)
	hash2, _ := HashPassword(password, DefaultParams)

	if hash1 == hash2 {
		t.Fatal("Two hashes of the same password should differ (different salts)")
	}
}

func TestVerifyPassword_CorrectPassword(t *testing.T) {
	password := "my-secure-password"

	hash, err := HashPassword(password, DefaultParams)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	match, err := VerifyPassword(password, hash)
	if err != nil {
		t.Fatalf("VerifyPassword returned error: %v", err)
	}

	if !match {
		t.Fatal("VerifyPassword should return true for correct password")
	}
}

func TestVerifyPassword_WrongPassword(t *testing.T) {
	password := "correct-password"
	wrongPassword := "wrong-password"

	hash, err := HashPassword(password, DefaultParams)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}

	match, err := VerifyPassword(wrongPassword, hash)
	if err != nil {
		t.Fatalf("VerifyPassword returned error: %v", err)
	}

	if match {
		t.Fatal("VerifyPassword should return false for wrong password")
	}
}

func TestVerifyPassword_InvalidHashFormat(t *testing.T) {
	tests := []struct {
		name string
		hash string
	}{
		{"empty hash", ""},
		{"too few parts", "$argon2id$v=19"},
		{"wrong algorithm", "$bcrypt$v=19$m=65536,t=3,p=2$c2FsdA==$aGFzaA=="},
		{"invalid version", "$argon2id$v=abc$m=65536,t=3,p=2$c2FsdA==$aGFzaA=="},
		{"invalid params", "$argon2id$v=19$invalid$params$c2FsdA==$aGFzaA=="},
		{"invalid salt base64", "$argon2id$v=19$m=65536,t=3,p=2$!!!invalid!!!$aGFzaA=="},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := VerifyPassword("any-password", tt.hash)
			if err == nil {
				t.Fatal("VerifyPassword should return error for invalid hash format")
			}
		})
	}
}

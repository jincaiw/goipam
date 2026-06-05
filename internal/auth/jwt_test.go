package auth

import (
	"testing"
	"time"

	"github.com/jincaiw/goipam/internal/config"
)

func newTestJWTManager() *JWTManager {
	return NewJWTManager(&config.SecurityConfig{
		JWTSecret:       "test-secret-key-for-testing",
		AccessTokenTTL:  30 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
	})
}

func TestGenerateAccessToken(t *testing.T) {
	mgr := newTestJWTManager()

	token, err := mgr.GenerateAccessToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateAccessToken returned error: %v", err)
	}

	if token == "" {
		t.Fatal("GenerateAccessToken returned empty token")
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	mgr := newTestJWTManager()

	token, err := mgr.GenerateRefreshToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateRefreshToken returned error: %v", err)
	}

	if token == "" {
		t.Fatal("GenerateRefreshToken returned empty token")
	}
}

func TestValidateAccessToken_Valid(t *testing.T) {
	mgr := newTestJWTManager()

	token, err := mgr.GenerateAccessToken(42, "testuser")
	if err != nil {
		t.Fatalf("GenerateAccessToken returned error: %v", err)
	}

	claims, err := mgr.ValidateAccessToken(token)
	if err != nil {
		t.Fatalf("ValidateAccessToken returned error: %v", err)
	}

	if claims.UserID != 42 {
		t.Fatalf("expected UserID=42, got %d", claims.UserID)
	}
	if claims.Username != "testuser" {
		t.Fatalf("expected Username=testuser, got %s", claims.Username)
	}
	if claims.Type != AccessToken {
		t.Fatalf("expected Type=AccessToken, got %s", claims.Type)
	}
}

func TestValidateAccessToken_RefreshTokenRejected(t *testing.T) {
	mgr := newTestJWTManager()

	refreshToken, err := mgr.GenerateRefreshToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateRefreshToken returned error: %v", err)
	}

	_, err = mgr.ValidateAccessToken(refreshToken)
	if err == nil {
		t.Fatal("ValidateAccessToken should reject refresh tokens")
	}
}

func TestValidateRefreshToken_Valid(t *testing.T) {
	mgr := newTestJWTManager()

	token, err := mgr.GenerateRefreshToken(42, "testuser")
	if err != nil {
		t.Fatalf("GenerateRefreshToken returned error: %v", err)
	}

	claims, err := mgr.ValidateRefreshToken(token)
	if err != nil {
		t.Fatalf("ValidateRefreshToken returned error: %v", err)
	}

	if claims.UserID != 42 {
		t.Fatalf("expected UserID=42, got %d", claims.UserID)
	}
	if claims.Type != RefreshToken {
		t.Fatalf("expected Type=RefreshToken, got %s", claims.Type)
	}
}

func TestValidateRefreshToken_AccessTokenRejected(t *testing.T) {
	mgr := newTestJWTManager()

	accessToken, err := mgr.GenerateAccessToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateAccessToken returned error: %v", err)
	}

	_, err = mgr.ValidateRefreshToken(accessToken)
	if err == nil {
		t.Fatal("ValidateRefreshToken should reject access tokens")
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
	mgr := NewJWTManager(&config.SecurityConfig{
		JWTSecret:       "test-secret-key",
		AccessTokenTTL:  -1 * time.Second, // Already expired
		RefreshTokenTTL: -1 * time.Second,
	})

	token, err := mgr.GenerateAccessToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateAccessToken returned error: %v", err)
	}

	_, err = mgr.ValidateAccessToken(token)
	if err == nil {
		t.Fatal("ValidateAccessToken should reject expired tokens")
	}
}

func TestValidateToken_WrongSecret(t *testing.T) {
	mgr1 := newTestJWTManager()
	mgr2 := NewJWTManager(&config.SecurityConfig{
		JWTSecret:       "different-secret-key",
		AccessTokenTTL:  30 * time.Minute,
		RefreshTokenTTL: 168 * time.Hour,
	})

	token, err := mgr1.GenerateAccessToken(1, "testuser")
	if err != nil {
		t.Fatalf("GenerateAccessToken returned error: %v", err)
	}

	_, err = mgr2.ValidateAccessToken(token)
	if err == nil {
		t.Fatal("ValidateAccessToken should reject tokens signed with a different secret")
	}
}

func TestValidateToken_InvalidTokenString(t *testing.T) {
	mgr := newTestJWTManager()

	_, err := mgr.ValidateAccessToken("not-a-valid-token")
	if err == nil {
		t.Fatal("ValidateAccessToken should reject invalid token strings")
	}
}

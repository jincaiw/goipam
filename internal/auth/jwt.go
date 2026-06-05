package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jincaiw/goipam/internal/config"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID   int64     `json:"user_id"`
	Username string    `json:"username"`
	Type     TokenType `json:"type"`
}

// JWTManager handles JWT token generation and validation.
type JWTManager struct {
	secret     []byte
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// NewJWTManager creates a new JWTManager.
func NewJWTManager(cfg *config.SecurityConfig) *JWTManager {
	return &JWTManager{
		secret:     []byte(cfg.JWTSecret),
		accessTTL:  cfg.AccessTokenTTL,
		refreshTTL: cfg.RefreshTokenTTL,
	}
}

// GenerateAccessToken generates a new access token.
func (m *JWTManager) GenerateAccessToken(userID int64, username string) (string, error) {
	now := time.Now()
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTTL)),
			Issuer:    "goipam",
		},
		UserID:   userID,
		Username: username,
		Type:     AccessToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

// GenerateRefreshToken generates a new refresh token.
func (m *JWTManager) GenerateRefreshToken(userID int64, username string) (string, error) {
	now := time.Now()
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(m.refreshTTL)),
			Issuer:    "goipam",
		},
		UserID:   userID,
		Username: username,
		Type:     RefreshToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secret)
}

// ValidateToken validates a token and returns its claims.
func (m *JWTManager) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

// ValidateAccessToken validates an access token specifically.
func (m *JWTManager) ValidateAccessToken(tokenStr string) (*Claims, error) {
	claims, err := m.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}
	if claims.Type != AccessToken {
		return nil, fmt.Errorf("not an access token")
	}
	return claims, nil
}

// ValidateRefreshToken validates a refresh token specifically.
func (m *JWTManager) ValidateRefreshToken(tokenStr string) (*Claims, error) {
	claims, err := m.ValidateToken(tokenStr)
	if err != nil {
		return nil, err
	}
	if claims.Type != RefreshToken {
		return nil, fmt.Errorf("not a refresh token")
	}
	return claims, nil
}

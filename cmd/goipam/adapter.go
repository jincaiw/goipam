package main

import (
	"github.com/jincaiw/goipam/internal/auth"
	"github.com/jincaiw/goipam/internal/rbac"
)

// authValidator adapts auth.AuthService to rbac.TokenValidator interface.
type authValidator struct {
	authService *auth.AuthService
}

func (v *authValidator) ValidateAccessToken(token string) (*rbac.Claims, error) {
	claims, err := v.authService.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}
	return &rbac.Claims{
		UserID:   claims.UserID,
		Username: claims.Username,
	}, nil
}

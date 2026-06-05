package rbac

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	userIDKey   contextKey = "user_id"
	usernameKey contextKey = "username"
)

// Claims represents JWT token claims.
type Claims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

// TokenValidator validates access tokens and returns claims.
type TokenValidator interface {
	ValidateAccessToken(token string) (*Claims, error)
}

// Middleware provides authentication and authorization middleware.
type Middleware struct {
	validator TokenValidator
}

// NewMiddleware creates a new RBAC Middleware.
func NewMiddleware(validator TokenValidator) *Middleware {
	return &Middleware{validator: validator}
}

// RequireAuth is middleware that requires a valid JWT token.
func (m *Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"success":false,"code":"AUTH_UNAUTHORIZED","message":"missing authorization token"}`, http.StatusUnauthorized)
			return
		}

		claims, err := m.validator.ValidateAccessToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"success":false,"code":"AUTH_INVALID_TOKEN","message":"invalid or expired token"}`, http.StatusUnauthorized)
			return
		}

		// Add user info to context
		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		ctx = context.WithValue(ctx, usernameKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID extracts the user ID from the request context.
func GetUserID(ctx context.Context) int64 {
	if v, ok := ctx.Value(userIDKey).(int64); ok {
		return v
	}
	return 0
}

// GetUsername extracts the username from the request context.
func GetUsername(ctx context.Context) string {
	if v, ok := ctx.Value(usernameKey).(string); ok {
		return v
	}
	return ""
}

// extractToken extracts the bearer token from the Authorization header.
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "bearer") {
		return ""
	}
	return strings.TrimSpace(parts[1])
}

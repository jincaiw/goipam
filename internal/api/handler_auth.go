package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/jincaiw/goipam/internal/api/errcode"
	"github.com/jincaiw/goipam/internal/auth"
	"github.com/jincaiw/goipam/internal/rbac"
)

// AuthHandler handles authentication API endpoints.
type AuthHandler struct {
	authService *auth.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login handles POST /api/v1/auth/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req auth.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, "VALIDATION_ERROR", "invalid request body")
		return
	}

	if req.Username == "" || req.Password == "" {
		BadRequest(w, r, "VALIDATION_ERROR", "username and password are required")
		return
	}

	if len(req.Username) > 128 {
		BadRequest(w, r, "VALIDATION_ERROR", "username is too long")
		return
	}
	if len(req.Password) > 128 {
		BadRequest(w, r, "VALIDATION_ERROR", "password is too long")
		return
	}

	resp, err := h.authService.Login(r.Context(), req)
	if err != nil {
		switch err {
		case auth.ErrInvalidCredentials:
			Error(w, r, http.StatusUnauthorized, errcode.AuthInvalidCredentials, "invalid username or password")
		case auth.ErrUserLocked:
			Error(w, r, http.StatusForbidden, errcode.AuthLocked, "account is temporarily locked")
		case auth.ErrUserDisabled:
			Error(w, r, http.StatusForbidden, errcode.AuthForbidden, "account is disabled")
		default:
			slog.Error("login failed", "error", err)
			InternalError(w, r, "login failed")
		}
		return
	}

	OK(w, r, resp)
}

// RefreshToken handles POST /api/v1/auth/refresh
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, "VALIDATION_ERROR", "invalid request body")
		return
	}

	if req.RefreshToken == "" {
		BadRequest(w, r, "VALIDATION_ERROR", "refresh_token is required")
		return
	}

	resp, err := h.authService.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		Error(w, r, http.StatusUnauthorized, errcode.AuthInvalidToken, "invalid or expired refresh token")
		return
	}

	OK(w, r, resp)
}

// Logout handles POST /api/v1/auth/logout
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// JWT is stateless, logout is handled client-side by discarding tokens
	OK(w, r, map[string]string{"message": "logged out"})
}

// Me handles GET /api/v1/auth/me
func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := rbac.GetUserID(r.Context())
	username := rbac.GetUsername(r.Context())

	OK(w, r, map[string]interface{}{
		"user_id":  userID,
		"username": username,
	})
}

// ChangePassword handles POST /api/v1/auth/change-password
func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		BadRequest(w, r, "VALIDATION_ERROR", "invalid request body")
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		BadRequest(w, r, "VALIDATION_ERROR", "old_password and new_password are required")
		return
	}

	if len(req.NewPassword) < 8 {
		BadRequest(w, r, "VALIDATION_ERROR", "new password must be at least 8 characters")
		return
	}
	if len(req.NewPassword) > 128 {
		BadRequest(w, r, "VALIDATION_ERROR", "new password must be at most 128 characters")
		return
	}

	userID := rbac.GetUserID(r.Context())
	if err := h.authService.ChangePassword(r.Context(), userID, req.OldPassword, req.NewPassword); err != nil {
		switch err {
		case auth.ErrInvalidCredentials:
			Error(w, r, http.StatusUnauthorized, errcode.AuthInvalidCredentials, "old password is incorrect")
		default:
			InternalError(w, r, "failed to change password")
		}
		return
	}

	OK(w, r, map[string]string{"message": "password changed"})
}

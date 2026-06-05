package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/jincaiw/goipam/internal/config"
	"github.com/jincaiw/goipam/internal/repository"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserLocked         = errors.New("user is locked")
	ErrUserDisabled       = errors.New("user is disabled")
	ErrUserNotFound       = errors.New("user not found")
)

const (
	maxFailedLogins = 5
	lockDuration    = 15 * time.Minute
)

// AuthService handles authentication operations.
type AuthService struct {
	userRepo *repository.UserRepository
	jwtMgr   *JWTManager
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo *repository.UserRepository, cfg *config.SecurityConfig) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtMgr:   NewJWTManager(cfg),
	}
}

// LoginRequest represents a login request.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents a login response.
type LoginResponse struct {
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	User         *repository.User `json:"user"`
}

// Login authenticates a user and returns tokens.
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("get user: %w", err)
	}

	// Check if user is disabled
	if user.Status == "disabled" {
		return nil, ErrUserDisabled
	}

	// Check if user is locked
	if user.LockedUntil != nil && user.LockedUntil.After(time.Now()) {
		return nil, ErrUserLocked
	}

	// Verify password
	match, err := VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("verify password: %w", err)
	}

	if !match {
		// Increment failed login count
		newCount := user.FailedLoginCount + 1
		var lockedUntil sql.NullTime
		if newCount >= maxFailedLogins {
			lockedUntil = sql.NullTime{
				Time:  time.Now().Add(lockDuration),
				Valid: true,
			}
			slog.Warn("user locked due to failed login attempts",
				"username", req.Username,
				"failed_count", newCount,
			)
		}
		if err := s.userRepo.UpdateLoginInfo(ctx, user.ID, newCount, lockedUntil); err != nil {
			slog.Error("failed to update login info", "error", err)
		}
		return nil, ErrInvalidCredentials
	}

	// Reset failed login count on successful login
	if err := s.userRepo.UpdateLoginInfo(ctx, user.ID, 0, sql.NullTime{}); err != nil {
		slog.Error("failed to reset login info", "error", err)
	}

	// Generate tokens
	accessToken, err := s.jwtMgr.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := s.jwtMgr.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         &user.User,
	}, nil
}

// RefreshToken refreshes an access token using a refresh token.
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*LoginResponse, error) {
	claims, err := s.jwtMgr.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Verify user still exists and is active
	user, err := s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, ErrUserNotFound
	}
	if user.Status == "disabled" {
		return nil, ErrUserDisabled
	}

	newAccessToken, err := s.jwtMgr.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("generate access token: %w", err)
	}

	newRefreshToken, err := s.jwtMgr.GenerateRefreshToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("generate refresh token: %w", err)
	}

	return &LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		User:         user,
	}, nil
}

// ChangePassword changes a user's password.
func (s *AuthService) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Get full user with password hash
	userFull, err := s.userRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		return ErrUserNotFound
	}

	match, err := VerifyPassword(oldPassword, userFull.PasswordHash)
	if err != nil {
		return fmt.Errorf("verify old password: %w", err)
	}
	if !match {
		return ErrInvalidCredentials
	}

	hash, err := HashPassword(newPassword, DefaultParams)
	if err != nil {
		return fmt.Errorf("hash new password: %w", err)
	}

	return s.userRepo.UpdatePassword(ctx, userID, hash)
}

// ValidateAccessToken validates an access token and returns user info.
func (s *AuthService) ValidateAccessToken(token string) (*Claims, error) {
	return s.jwtMgr.ValidateAccessToken(token)
}

// JWTManager returns the JWT manager for middleware use.
func (s *AuthService) GetJWTManager() *JWTManager {
	return s.jwtMgr
}

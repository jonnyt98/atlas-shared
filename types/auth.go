package types

import (
	"time"
)

// AuthTokens represents JWT authentication tokens
type AuthTokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
}

// AuthResponse represents the complete authentication response
type AuthResponse struct {
	AuthTokens
	User UserResponse `json:"user"`
}

// LoginRequest represents user login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RegisterRequest represents user registration request
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RefreshTokenRequest represents token refresh request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshToken represents a stored refresh token
type RefreshToken struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
}

// GoogleOAuthRequest represents Google OAuth request
type GoogleOAuthRequest struct {
	Code  string `json:"code" validate:"required"`
	State string `json:"state,omitempty"`
}

// GoogleOAuthResponse represents Google OAuth response
type GoogleOAuthResponse struct {
	AuthResponse
	IsNewUser bool `json:"is_new_user"`
}

// LogoutRequest represents logout request
type LogoutRequest struct {
	RefreshToken string `json:"refresh_token,omitempty"`
}

// ForgotPasswordRequest represents forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// ResetPasswordRequest represents reset password request
type ResetPasswordRequest struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

// ChangePasswordRequest represents change password request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=8"`
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	UserID    string `json:"user_id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"` // "access" or "refresh"
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
}

// SessionInfo represents current session information
type SessionInfo struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// EmailVerificationSendRequest represents email verification send request
type EmailVerificationSendRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

// AuthError represents authentication-specific errors
type AuthError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Common auth error codes
const (
	AuthErrorInvalidCredentials = "invalid_credentials"
	AuthErrorUserNotFound      = "user_not_found"
	AuthErrorUserAlreadyExists = "user_already_exists"
	AuthErrorInvalidToken      = "invalid_token"
	AuthErrorTokenExpired      = "token_expired"
	AuthErrorEmailNotVerified  = "email_not_verified"
	AuthErrorAccountLocked     = "account_locked"
	AuthErrorInvalidGoogleCode = "invalid_google_code"
)
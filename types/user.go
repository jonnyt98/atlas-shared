package types

import (
	"time"

	"github.com/google/uuid"
)

// User represents the core user entity shared across services
type User struct {
	ID                    string     `json:"id"`
	Email                 string     `json:"email"`
	PasswordHash          *string    `json:"-"`
	GoogleID              *string    `json:"-"`
	Role                  string     `json:"role"`
	OrgID                 *uuid.UUID `json:"org_id,omitempty"`
	EmailVerified         bool       `json:"email_verified"`
	HasActiveSubscription bool       `json:"has_active_subscription"`
	StripeCustomerID      *string    `json:"-"`
	StripeSubscriptionID  *string    `json:"-"`
	SubscriptionTier      *string    `json:"subscription_tier,omitempty"`
	SubscriptionStatus    *string    `json:"subscription_status,omitempty"`
	VerificationToken     *string    `json:"-"`
	TokenExpiresAt        *time.Time `json:"-"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeletedAt             *time.Time `json:"deleted_at,omitempty"`
}

// UserCreateRequest represents the request to create a new user
type UserCreateRequest struct {
	Email             string     `json:"email" validate:"required,email"`
	PasswordHash      string     `json:"password_hash,omitempty"`
	Role              string     `json:"role,omitempty"`
	OrgID             *uuid.UUID `json:"org_id,omitempty"`
	VerificationToken string     `json:"verification_token,omitempty"`
	TokenExpiresAt    *time.Time `json:"token_expires_at,omitempty"`
}

// UserUpdateRequest represents the request to update user data
type UserUpdateRequest struct {
	Email  *string    `json:"email,omitempty"`
	Role   *string    `json:"role,omitempty"`
	OrgID  *uuid.UUID `json:"org_id,omitempty"`
}

// UserResponse represents the public user data returned to clients
type UserResponse struct {
	ID                    string     `json:"id"`
	Email                 string     `json:"email"`
	Role                  string     `json:"role"`
	OrgID                 *uuid.UUID `json:"org_id,omitempty"`
	EmailVerified         bool       `json:"email_verified"`
	HasActiveSubscription bool       `json:"has_active_subscription"`
	SubscriptionTier      *string    `json:"subscription_tier,omitempty"`
	SubscriptionStatus    *string    `json:"subscription_status,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
}

// ToResponse converts a User to UserResponse, excluding sensitive fields
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:                    u.ID,
		Email:                 u.Email,
		Role:                  u.Role,
		OrgID:                 u.OrgID,
		EmailVerified:         u.EmailVerified,
		HasActiveSubscription: u.HasActiveSubscription,
		SubscriptionTier:      u.SubscriptionTier,
		SubscriptionStatus:    u.SubscriptionStatus,
		CreatedAt:             u.CreatedAt,
	}
}

// EmailVerificationRequest represents email verification update request
type EmailVerificationRequest struct {
	EmailVerified     *bool      `json:"email_verified,omitempty"`
	VerificationToken *string    `json:"verification_token,omitempty"`
	TokenExpiresAt    *time.Time `json:"token_expires_at,omitempty"`
}

// PasswordUpdateRequest represents password update request
type PasswordUpdateRequest struct {
	PasswordHash string `json:"password_hash" validate:"required"`
}

// GoogleLinkRequest represents Google account linking request
type GoogleLinkRequest struct {
	GoogleID string `json:"google_id" validate:"required"`
}

// AuthenticateRequest represents user authentication request
type AuthenticateRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthenticateResponse represents user authentication response
type AuthenticateResponse struct {
	User UserResponse `json:"user"`
}

// VerifyEmailRequest represents email verification request
type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required"`
}

// VerifyEmailResponse represents email verification response
type VerifyEmailResponse struct {
	Message string       `json:"message"`
	User    UserResponse `json:"user"`
}

// UserListQuery represents query parameters for listing users
type UserListQuery struct {
	Page          int        `json:"page,omitempty"`
	Limit         int        `json:"limit,omitempty"`
	Role          string     `json:"role,omitempty"`
	OrgID         *uuid.UUID `json:"org_id,omitempty"`
	EmailVerified *bool      `json:"email_verified,omitempty"`
}

// UserListResponse represents the response for user listing
type UserListResponse struct {
	Users      []UserResponse `json:"users"`
	Total      int            `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}
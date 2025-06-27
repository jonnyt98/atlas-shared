package contracts

import (
	"context"

	"github.com/jonnyt98/atlas-shared/types"
	"github.com/google/uuid"
)

// UserServiceClient defines the interface for interacting with the user service
type UserServiceClient interface {
	// User CRUD operations
	CreateUser(ctx context.Context, req types.UserCreateRequest) (*types.UserResponse, error)
	GetUserByID(ctx context.Context, id string) (*types.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*types.UserResponse, error)
	GetUserByGoogleID(ctx context.Context, googleID string) (*types.UserResponse, error)
	UpdateUser(ctx context.Context, id string, req types.UserUpdateRequest) (*types.UserResponse, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, query types.UserListQuery) (*types.UserListResponse, error)

	// Email verification
	UpdateEmailVerification(ctx context.Context, id string, req types.EmailVerificationRequest) (*types.UserResponse, error)
	VerifyEmail(ctx context.Context, req types.VerifyEmailRequest) (*types.VerifyEmailResponse, error)
	ResendVerification(ctx context.Context, id string) error

	// Authentication
	AuthenticateUser(ctx context.Context, req types.AuthenticateRequest) (*types.AuthenticateResponse, error)

	// Google OAuth
	LinkGoogleAccount(ctx context.Context, id string, req types.GoogleLinkRequest) (*types.UserResponse, error)

	// Password management
	UpdatePassword(ctx context.Context, id string, req types.PasswordUpdateRequest) (*types.UserResponse, error)

	// Subscription management
	UpdateSubscription(ctx context.Context, id string, req types.SubscriptionUpdateRequest) (*types.UserResponse, error)
	GetSubscription(ctx context.Context, id string) (*types.UserSubscriptionResponse, error)
	CancelSubscription(ctx context.Context, id string) (*types.UserResponse, error)

	// User existence check
	UserExists(ctx context.Context, email string) (bool, error)

	// Health check
	Health(ctx context.Context) error
}

// OrganizationServiceClient defines the interface for organization operations
type OrganizationServiceClient interface {
	// Organization CRUD operations
	CreateOrganization(ctx context.Context, req types.OrganizationCreateRequest) (*types.OrganizationResponse, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*types.OrganizationResponse, error)
	UpdateOrganization(ctx context.Context, id uuid.UUID, req types.OrganizationUpdateRequest) (*types.OrganizationResponse, error)
	DeleteOrganization(ctx context.Context, id uuid.UUID) error
	ListOrganizations(ctx context.Context, query types.OrganizationListQuery) (*types.OrganizationListResponse, error)

	// Organization users
	ListOrganizationUsers(ctx context.Context, id uuid.UUID, query types.UserListQuery) (*types.UserListResponse, error)
	AddUserToOrganization(ctx context.Context, orgID uuid.UUID, req types.AddUserToOrgRequest) (*types.UserResponse, error)
	RemoveUserFromOrganization(ctx context.Context, orgID uuid.UUID, userID string) error
}

// Organization types (adding to complete the contract)
type OrganizationCreateRequest struct {
	Name     string                 `json:"name" validate:"required"`
	Slug     string                 `json:"slug" validate:"required"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

type OrganizationUpdateRequest struct {
	Name     *string                `json:"name,omitempty"`
	Slug     *string                `json:"slug,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

type OrganizationResponse struct {
	ID        uuid.UUID              `json:"id"`
	Name      string                 `json:"name"`
	Slug      string                 `json:"slug"`
	Settings  map[string]interface{} `json:"settings"`
	CreatedAt string                 `json:"created_at"`
}

type OrganizationListQuery struct {
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Name  string `json:"name,omitempty"`
}

type OrganizationListResponse struct {
	Organizations []OrganizationResponse `json:"organizations"`
	Total         int                    `json:"total"`
	Page          int                    `json:"page"`
	Limit         int                    `json:"limit"`
	TotalPages    int                    `json:"total_pages"`
}

type AddUserToOrgRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role,omitempty"`
}
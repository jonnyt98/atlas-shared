package types

import (
	"time"

	"github.com/google/uuid"
)

// Organization represents the core organization entity
type Organization struct {
	ID        uuid.UUID              `json:"id"`
	Name      string                 `json:"name"`
	Slug      string                 `json:"slug"`
	Settings  map[string]interface{} `json:"settings"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty"`
}

// OrganizationCreateRequest represents the request to create a new organization
type OrganizationCreateRequest struct {
	Name     string                 `json:"name" validate:"required"`
	Slug     string                 `json:"slug" validate:"required"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// OrganizationUpdateRequest represents the request to update organization data
type OrganizationUpdateRequest struct {
	Name     *string                `json:"name,omitempty"`
	Slug     *string                `json:"slug,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"`
}

// OrganizationResponse represents the public organization data
type OrganizationResponse struct {
	ID        uuid.UUID              `json:"id"`
	Name      string                 `json:"name"`
	Slug      string                 `json:"slug"`
	Settings  map[string]interface{} `json:"settings"`
	CreatedAt time.Time              `json:"created_at"`
}

// ToResponse converts an Organization to OrganizationResponse
func (o *Organization) ToResponse() OrganizationResponse {
	return OrganizationResponse{
		ID:        o.ID,
		Name:      o.Name,
		Slug:      o.Slug,
		Settings:  o.Settings,
		CreatedAt: o.CreatedAt,
	}
}

// OrganizationListQuery represents query parameters for listing organizations
type OrganizationListQuery struct {
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
	Name  string `json:"name,omitempty"`
}

// OrganizationListResponse represents the response for organization listing
type OrganizationListResponse struct {
	Organizations []OrganizationResponse `json:"organizations"`
	Total         int                    `json:"total"`
	Page          int                    `json:"page"`
	Limit         int                    `json:"limit"`
	TotalPages    int                    `json:"total_pages"`
}

// AddUserToOrgRequest represents adding a user to an organization
type AddUserToOrgRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Role   string `json:"role,omitempty"`
}
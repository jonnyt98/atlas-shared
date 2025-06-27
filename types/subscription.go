package types

import (
	"time"
)

// Subscription represents the core subscription entity
type Subscription struct {
	ID                 string     `json:"id"`
	UserID             string     `json:"user_id"`
	StripeCustomerID   string     `json:"stripe_customer_id"`
	StripeSubscriptionID string   `json:"stripe_subscription_id"`
	Plan               string     `json:"plan"`
	Tier               string     `json:"tier"`
	Status             string     `json:"status"`
	StartedAt          time.Time  `json:"started_at"`
	CurrentPeriodStart time.Time  `json:"current_period_start"`
	CurrentPeriodEnd   time.Time  `json:"current_period_end"`
	CanceledAt         *time.Time `json:"canceled_at,omitempty"`
	TrialEndsAt        *time.Time `json:"trial_ends_at,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

// SubscriptionStatus represents valid subscription statuses
type SubscriptionStatus string

const (
	SubscriptionStatusActive    SubscriptionStatus = "active"
	SubscriptionStatusCanceled  SubscriptionStatus = "canceled"
	SubscriptionStatusPastDue   SubscriptionStatus = "past_due"
	SubscriptionStatusTrialing  SubscriptionStatus = "trialing"
	SubscriptionStatusIncomplete SubscriptionStatus = "incomplete"
)

// SubscriptionTier represents valid subscription tiers
type SubscriptionTier string

const (
	SubscriptionTierBasic      SubscriptionTier = "basic"
	SubscriptionTierPro        SubscriptionTier = "pro"
	SubscriptionTierEnterprise SubscriptionTier = "enterprise"
)

// CheckoutSessionRequest represents the request for creating a checkout session
type CheckoutSessionRequest struct {
	CustomerID string `json:"customerId" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Tier       string `json:"tier" validate:"required"`
}

// CheckoutSessionResponse represents the response from creating a checkout session
type CheckoutSessionResponse struct {
	SessionID string `json:"sessionId"`
}

// SubscriptionCreateRequest represents the request to create a subscription
type SubscriptionCreateRequest struct {
	UserID               string                `json:"user_id" validate:"required"`
	StripeCustomerID     string                `json:"stripe_customer_id" validate:"required"`
	StripeSubscriptionID string                `json:"stripe_subscription_id" validate:"required"`
	Plan                 string                `json:"plan" validate:"required"`
	Tier                 SubscriptionTier      `json:"tier" validate:"required"`
	Status               SubscriptionStatus    `json:"status" validate:"required"`
	StartedAt            time.Time             `json:"started_at"`
	CurrentPeriodStart   time.Time             `json:"current_period_start"`
	CurrentPeriodEnd     time.Time             `json:"current_period_end"`
	TrialEndsAt          *time.Time            `json:"trial_ends_at,omitempty"`
}

// SubscriptionUpdateRequest represents the request to update a subscription
type SubscriptionUpdateRequest struct {
	Status               *SubscriptionStatus `json:"status,omitempty"`
	Tier                 *SubscriptionTier   `json:"tier,omitempty"`
	Plan                 *string             `json:"plan,omitempty"`
	StripeCustomerID     *string             `json:"stripe_customer_id,omitempty"`
	HasActiveSubscription *bool              `json:"has_active_subscription,omitempty"`
	CurrentPeriodStart   *time.Time          `json:"current_period_start,omitempty"`
	CurrentPeriodEnd     *time.Time          `json:"current_period_end,omitempty"`
	CanceledAt           *time.Time          `json:"canceled_at,omitempty"`
	TrialEndsAt          *time.Time          `json:"trial_ends_at,omitempty"`
}

// SubscriptionResponse represents the public subscription data
type SubscriptionResponse struct {
	ID                 string             `json:"id"`
	UserID             string             `json:"user_id"`
	Plan               string             `json:"plan"`
	Tier               SubscriptionTier   `json:"tier"`
	Status             SubscriptionStatus `json:"status"`
	StartedAt          time.Time          `json:"started_at"`
	CurrentPeriodStart time.Time          `json:"current_period_start"`
	CurrentPeriodEnd   time.Time          `json:"current_period_end"`
	CanceledAt         *time.Time         `json:"canceled_at,omitempty"`
	TrialEndsAt        *time.Time         `json:"trial_ends_at,omitempty"`
	CreatedAt          time.Time          `json:"created_at"`
}

// ToResponse converts a Subscription to SubscriptionResponse, excluding sensitive fields
func (s *Subscription) ToResponse() SubscriptionResponse {
	return SubscriptionResponse{
		ID:                 s.ID,
		UserID:             s.UserID,
		Plan:               s.Plan,
		Tier:               SubscriptionTier(s.Tier),
		Status:             SubscriptionStatus(s.Status),
		StartedAt:          s.StartedAt,
		CurrentPeriodStart: s.CurrentPeriodStart,
		CurrentPeriodEnd:   s.CurrentPeriodEnd,
		CanceledAt:         s.CanceledAt,
		TrialEndsAt:        s.TrialEndsAt,
		CreatedAt:          s.CreatedAt,
	}
}

// UserSubscriptionResponse represents subscription data from user service perspective
type UserSubscriptionResponse struct {
	UserID       string             `json:"user_id"`
	Subscription *SubscriptionResponse `json:"subscription,omitempty"`
}

// StripeWebhookEvent represents incoming Stripe webhook events
type StripeWebhookEvent struct {
	Type   string      `json:"type"`
	Data   interface{} `json:"data"`
	Object string      `json:"object"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}
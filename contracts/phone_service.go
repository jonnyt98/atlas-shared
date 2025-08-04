package contracts

import (
	"context"
	"time"
	"github.com/jonnyt98/atlas-shared/types"
	"github.com/google/uuid"
)

// PhoneServiceClient defines the interface for phone service operations
type PhoneServiceClient interface {
	// Phone number management
	ProvisionPhoneNumber(ctx context.Context, req types.PhoneProvisionRequest) (*types.PhoneProvisionResponse, error)
	ReleasePhoneNumber(ctx context.Context, phoneNumberID uuid.UUID) error
	GetUserPhoneNumbers(ctx context.Context, userID string) ([]types.PhoneNumber, error)
	GetPhoneNumberUsage(ctx context.Context, phoneNumberID uuid.UUID) ([]types.PhoneUsage, error)
	
	// Phone number configuration
	UpdatePhoneNumberConfiguration(ctx context.Context, phoneNumberID uuid.UUID, config map[string]interface{}) error
	SuspendPhoneNumber(ctx context.Context, phoneNumberID uuid.UUID) error
	ReactivatePhoneNumber(ctx context.Context, phoneNumberID uuid.UUID) error
	
	// Analytics and reporting
	GetPhoneNumberAnalytics(ctx context.Context, userID string, from, to time.Time) (*types.PhoneUsageAnalytics, error)
	GetServiceHealth(ctx context.Context) (*types.HealthResponse, error)
}

// PhoneProviderServiceClient defines the interface for simple phone provider operations (matches actual service)
type PhoneProviderServiceClient interface {
	// Simple phone number purchase
	PurchasePhoneNumber(ctx context.Context, req types.PurchasePhoneNumberRequest) (*types.PurchasePhoneNumberResponse, error)
	GetServiceHealth(ctx context.Context) (*types.HealthResponse, error)
}

// PhoneWebhookHandler defines the interface for handling Twilio webhooks
type PhoneWebhookHandler interface {
	HandleVoiceWebhook(ctx context.Context, data map[string]string) error
	HandleSMSWebhook(ctx context.Context, data map[string]string) error
	HandleStatusWebhook(ctx context.Context, data map[string]string) error
}
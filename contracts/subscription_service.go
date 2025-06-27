package contracts

import (
	"context"

	"github.com/jonnyt98/atlas-shared/types"
)

// SubscriptionServiceClient defines the interface for interacting with the subscription service
type SubscriptionServiceClient interface {
	// Checkout and payment
	CreateCheckoutSession(ctx context.Context, req types.CheckoutSessionRequest) (*types.CheckoutSessionResponse, error)
	
	// Subscription CRUD operations
	CreateSubscription(ctx context.Context, req types.SubscriptionCreateRequest) (*types.SubscriptionResponse, error)
	GetSubscriptionByStripeID(ctx context.Context, stripeSubscriptionID string) (*types.SubscriptionResponse, error)
	GetSubscriptionByUserID(ctx context.Context, userID string) (*types.SubscriptionResponse, error)
	UpdateSubscriptionByStripeID(ctx context.Context, stripeSubscriptionID string, req types.SubscriptionUpdateRequest) (*types.SubscriptionResponse, error)
	CancelSubscription(ctx context.Context, subscriptionID string) (*types.SubscriptionResponse, error)
	
	// Stripe webhook handling
	HandleStripeWebhook(ctx context.Context, payload []byte, signature string) error
	
	// Health check
	Health(ctx context.Context) error
}

// AuthServiceClient defines the interface for interacting with the auth service
type AuthServiceClient interface {
	// Authentication
	Login(ctx context.Context, req types.LoginRequest) (*types.AuthResponse, error)
	Register(ctx context.Context, req types.RegisterRequest) (*types.AuthResponse, error)
	RefreshToken(ctx context.Context, req types.RefreshTokenRequest) (*types.AuthResponse, error)
	Logout(ctx context.Context, req types.LogoutRequest) error
	
	// Google OAuth
	GoogleOAuth(ctx context.Context, req types.GoogleOAuthRequest) (*types.GoogleOAuthResponse, error)
	GetGoogleOAuthURL(ctx context.Context, state string) (string, error)
	
	// Password management
	ForgotPassword(ctx context.Context, req types.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, req types.ResetPasswordRequest) error
	ChangePassword(ctx context.Context, userID string, req types.ChangePasswordRequest) error
	
	// Token validation and session
	ValidateToken(ctx context.Context, token string) (*types.SessionInfo, error)
	GetSessionInfo(ctx context.Context, token string) (*types.SessionInfo, error)
	
	// Email verification
	SendEmailVerification(ctx context.Context, req types.EmailVerificationSendRequest) error
	
	// Health check
	Health(ctx context.Context) error
}

// EmailServiceClient defines the interface for interacting with the email service
type EmailServiceClient interface {
	// Email sending
	SendVerificationEmail(ctx context.Context, userID, email, token string) error
	SendWelcomeEmail(ctx context.Context, userID, email, name string) error
	SendPasswordResetEmail(ctx context.Context, userID, email, token string) error
	SendSubscriptionConfirmationEmail(ctx context.Context, userID, email string, subscription types.SubscriptionResponse) error
	SendSubscriptionCancelationEmail(ctx context.Context, userID, email string) error
	
	// Health check
	Health(ctx context.Context) error
}
package types

import (
	"time"
	"github.com/google/uuid"
)

// PhoneNumber represents a phone number entity
type PhoneNumber struct {
	ID            uuid.UUID              `json:"id"`
	UserID        string                 `json:"user_id"`
	Number        string                 `json:"number"`
	TwilioSID     string                 `json:"twilio_sid"`
	Status        string                 `json:"status"`
	Capabilities  []string               `json:"capabilities"`
	AreaCode      string                 `json:"area_code,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
}

// PhoneUsage represents phone usage analytics
type PhoneUsage struct {
	ID              uuid.UUID              `json:"id"`
	PhoneNumberID   uuid.UUID              `json:"phone_number_id"`
	UsageType       string                 `json:"usage_type"`
	FromNumber      string                 `json:"from_number,omitempty"`
	ToNumber        string                 `json:"to_number,omitempty"`
	DurationSeconds int                    `json:"duration_seconds,omitempty"`
	CostCents       int                    `json:"cost_cents,omitempty"`
	TwilioSID       string                 `json:"twilio_sid,omitempty"`
	Status          string                 `json:"status"`
	Metadata        map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt       time.Time              `json:"created_at"`
}

// PhoneProvisionRequest represents a request to provision a new phone number
type PhoneProvisionRequest struct {
	UserID       string   `json:"user_id" binding:"required"`
	AreaCode     string   `json:"area_code,omitempty"`
	Capabilities []string `json:"capabilities" binding:"required"`
}

// PhoneProvisionResponse represents the response after provisioning a phone number
type PhoneProvisionResponse struct {
	ID           uuid.UUID `json:"id"`
	UserID       string    `json:"user_id"`
	Number       string    `json:"number"`
	TwilioSID    string    `json:"twilio_sid"`
	Status       string    `json:"status"`
	Capabilities []string  `json:"capabilities"`
	AreaCode     string    `json:"area_code,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// PhoneNumberListResponse represents a list of phone numbers
type PhoneNumberListResponse struct {
	PhoneNumbers []PhoneNumber `json:"phone_numbers"`
}

// PhoneUsageListResponse represents a list of phone usage records
type PhoneUsageListResponse struct {
	Usage []PhoneUsage `json:"usage"`
}

// PhoneNumberStatus constants
const (
	PhoneNumberStatusActive    = "active"
	PhoneNumberStatusInactive  = "inactive"
	PhoneNumberStatusSuspended = "suspended"
	PhoneNumberStatusReleased  = "released"
)

// PhoneUsageType constants
const (
	PhoneUsageTypeVoiceInbound  = "voice_inbound"
	PhoneUsageTypeVoiceOutbound = "voice_outbound"
	PhoneUsageTypeSMSInbound    = "sms_inbound"
	PhoneUsageTypeSMSOutbound   = "sms_outbound"
	PhoneUsageTypeMMSInbound    = "mms_inbound"
	PhoneUsageTypeMMSOutbound   = "mms_outbound"
)

// PhoneCapability constants
const (
	PhoneCapabilityVoice = "voice"
	PhoneCapabilitySMS   = "sms"
	PhoneCapabilityMMS   = "mms"
)

// PurchasePhoneNumberRequest represents a request to purchase a new phone number (simplified API)
type PurchasePhoneNumberRequest struct {
	AreaCode    string `json:"area_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// PurchasePhoneNumberResponse represents the response after purchasing a phone number (simplified API)
type PurchasePhoneNumberResponse struct {
	PhoneNumber string `json:"phone_number"`
	Sid         string `json:"sid"`
	Status      string `json:"status"`
}

// PhoneUsageAnalytics represents aggregated phone usage analytics
type PhoneUsageAnalytics struct {
	UserID           string                 `json:"user_id"`
	TotalCalls       int                    `json:"total_calls"`
	TotalSMS         int                    `json:"total_sms"`
	TotalCostCents   int                    `json:"total_cost_cents"`
	UsageByType      map[string]int         `json:"usage_by_type"`
	UsageByDay       map[string]int         `json:"usage_by_day"`
	TopNumbers       []string               `json:"top_numbers"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	PeriodStart      time.Time              `json:"period_start"`
	PeriodEnd        time.Time              `json:"period_end"`
}

// Phone service status constants
const (
	PhoneNumberStatusPurchased = "purchased"
	DefaultCountryCode        = "US"
)
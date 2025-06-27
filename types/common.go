package types

import "time"

// APIResponse represents a standard API response wrapper
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

// APIError represents a standard API error
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Meta represents metadata for paginated responses
type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Version   string    `json:"version,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Services  map[string]ServiceHealth `json:"services,omitempty"`
}

// ServiceHealth represents individual service health status
type ServiceHealth struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Latency string `json:"latency,omitempty"`
}

// Common HTTP status responses
const (
	StatusOK       = "ok"
	StatusError    = "error"
	StatusHealthy  = "healthy"
	StatusDegraded = "degraded"
	StatusDown     = "down"
)

// Common error codes
const (
	ErrorCodeValidation      = "validation_error"
	ErrorCodeNotFound        = "not_found"
	ErrorCodeUnauthorized    = "unauthorized"
	ErrorCodeForbidden       = "forbidden"
	ErrorCodeConflict        = "conflict"
	ErrorCodeInternalServer  = "internal_server_error"
	ErrorCodeBadRequest      = "bad_request"
	ErrorCodeServiceUnavailable = "service_unavailable"
)

// PaginationDefaults
const (
	DefaultPage  = 1
	DefaultLimit = 20
	MaxLimit     = 100
)
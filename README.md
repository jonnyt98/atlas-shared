# Atlas Shared Types and Contracts

This module contains shared types and service contracts for the Atlas microservices architecture.

## Structure

```
shared/
├── types/           # Shared data types
│   ├── user.go     # User-related types
│   ├── auth.go     # Authentication types
│   ├── subscription.go # Subscription types
│   ├── organization.go # Organization types
│   └── common.go   # Common API types
├── contracts/      # Service interface contracts
│   ├── user_service.go       # User service interface
│   └── subscription_service.go # Subscription/Auth/Email service interfaces
├── go.mod
└── README.md
```

## Usage

### In your service's go.mod

```go
require (
    github.com/jonnyt98/atlas-shared v0.1.0
)
```

### Import types

```go
import (
    "github.com/jonnyt98/atlas-shared/types"
    "github.com/jonnyt98/atlas-shared/contracts"
)
```

### Example usage

```go
// Using shared types
func CreateUser(req types.UserCreateRequest) (*types.UserResponse, error) {
    // Implementation
}

// Implementing service contracts
type UserService struct {}

func (s *UserService) CreateUser(ctx context.Context, req types.UserCreateRequest) (*types.UserResponse, error) {
    // Implementation
}
```

## Types Overview

### User Types
- `User`: Core user entity
- `UserCreateRequest/UpdateRequest`: CRUD operations
- `UserResponse`: Public user data
- `AuthenticateRequest/Response`: Authentication flow

### Auth Types
- `AuthTokens`: JWT token structure
- `AuthResponse`: Complete auth response with user data
- `LoginRequest/RegisterRequest`: Authentication requests
- `RefreshTokenRequest`: Token refresh flow

### Subscription Types
- `Subscription`: Core subscription entity
- `CheckoutSessionRequest/Response`: Stripe checkout flow
- `SubscriptionUpdateRequest`: Subscription modifications
- `SubscriptionStatus/Tier`: Enums for subscription states

### Organization Types
- `Organization`: Core organization entity
- `OrganizationCreateRequest/UpdateRequest`: CRUD operations
- `AddUserToOrgRequest`: Organization membership

### Common Types
- `APIResponse`: Standard API response wrapper
- `APIError`: Standard error format
- `HealthResponse`: Health check format
- Error codes and pagination constants

## Service Contracts

### UserServiceClient
Interface for user service operations including CRUD, authentication, email verification, and subscription management.

### SubscriptionServiceClient
Interface for subscription service operations including Stripe integration and webhook handling.

### AuthServiceClient
Interface for authentication service operations including login, registration, OAuth, and token management.

### EmailServiceClient
Interface for email service operations including verification emails, welcome emails, and subscription notifications.

## Migration Guide

When migrating existing services to use shared types:

1. Add the shared module dependency
2. Replace local type definitions with shared types
3. Update import statements
4. Implement the service contract interfaces
5. Update API handlers to use shared request/response types

## Versioning

This module follows semantic versioning. Breaking changes will increment the major version.
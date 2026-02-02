package ports

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
)

// IdentityClient defines the interface for interacting with Identity Service
type IdentityClient interface {
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error)
	DeleteUser(ctx context.Context, userID int64) error
	UpdateTenantPlan(ctx context.Context, tenantID int64, planID int64) error
}

// BillingClient defines the interface for interacting with Billing Service
type BillingClient interface {
	// Subscription
	CreateSubscription(ctx context.Context, userID int64, planID int) (int64, error)
	GetSubscription(ctx context.Context, subscriptionID int64) (*dto.SubscriptionResponse, error)
	CancelSubscription(ctx context.Context, subscriptionID int64) error
	UpdateSubscription(ctx context.Context, tenantID int64, req dto.UpdateSubscriptionRequest) error

	// Plan
	GetPlan(ctx context.Context, planID int) (*dto.PlanResponse, error)

	// Invoice
	CreateInvoice(ctx context.Context, req dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error)
	UpdateInvoice(ctx context.Context, invoiceID int64, req dto.UpdateInvoiceRequest) error
}

// PaymentClient defines the interface for interacting with Payment Service
type PaymentClient interface {
	ProcessPayment(ctx context.Context, req dto.ProcessPaymentRequest) (*dto.PaymentResponse, error)
}

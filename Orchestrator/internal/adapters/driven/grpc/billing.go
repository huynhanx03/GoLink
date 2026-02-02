package grpc

import (
	"context"

	billingv1 "go-link/common/gen/go/billing/v1"
	"go-link/common/pkg/grpc"
	"go-link/common/pkg/settings"
	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/ports"
)

type billingClientAdapter struct {
	client billingv1.BillingServiceClient
}

func NewBillingClient(cfg settings.GRPCService) (ports.BillingClient, error) {
	conn, err := grpc.NewClientConn(cfg)
	if err != nil {
		return nil, err
	}
	return &billingClientAdapter{
		client: billingv1.NewBillingServiceClient(conn),
	}, nil
}

// -- Existing Methods (Real gRPC) --

func (a *billingClientAdapter) CreateSubscription(ctx context.Context, userID int64, planID int) (int64, error) {
	resp, err := a.client.CreateSubscription(ctx, &billingv1.CreateSubscriptionRequest{
		UserId: userID,
		PlanId: int64(planID),
	})
	if err != nil {
		return 0, err
	}
	return resp.SubscriptionId, nil
}

func (a *billingClientAdapter) CancelSubscription(ctx context.Context, subscriptionID int64) error {
	_, err := a.client.CancelSubscription(ctx, &billingv1.CancelSubscriptionRequest{
		SubscriptionId: subscriptionID,
	})
	return err
}

// -- New Methods for Upgrade SAGA (Currently Stubbed) --
// TODO: Update proto definitions and implement real gRPC calls

func (a *billingClientAdapter) GetSubscription(ctx context.Context, subscriptionID int64) (*dto.SubscriptionResponse, error) {
	// Stub implementation
	return &dto.SubscriptionResponse{
		ID:       subscriptionID,
		TenantID: 1001,
		PlanID:   1, // Free plan
		Status:   "ACTIVE",
	}, nil
}

func (a *billingClientAdapter) UpdateSubscription(ctx context.Context, tenantID int64, req dto.UpdateSubscriptionRequest) error {
	// Stub implementation
	return nil
}

func (a *billingClientAdapter) GetPlan(ctx context.Context, planID int) (*dto.PlanResponse, error) {
	// Stub implementation
	return &dto.PlanResponse{
		ID:        planID,
		Name:      "Pro Plan",
		BasePrice: 20.00,
		Period:    "month",
	}, nil
}

func (a *billingClientAdapter) CreateInvoice(ctx context.Context, req dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	// Stub implementation
	return &dto.InvoiceResponse{
		ID:             67890,
		SubscriptionID: req.SubscriptionID,
		TenantID:       req.TenantID,
		Amount:         req.Amount,
		Currency:       req.Currency,
		Status:         req.Status,
	}, nil
}

func (a *billingClientAdapter) UpdateInvoice(ctx context.Context, invoiceID int64, req dto.UpdateInvoiceRequest) error {
	// Stub implementation
	return nil
}

package grpc

import (
	"context"

	billingv1 "go-link/common/gen/go/billing/v1"
	"go-link/common/pkg/grpc"
	"go-link/common/pkg/settings"
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

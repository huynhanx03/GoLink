package service

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/core/saga/change_plan"
	"go-link/orchestrator/internal/core/saga/registration"
	"go-link/orchestrator/internal/ports"
)

type orchestratorService struct {
	identityClient ports.IdentityClient
	billingClient  ports.BillingClient
	paymentClient  ports.PaymentClient
}

func NewOrchestratorService(
	identityClient ports.IdentityClient,
	billingClient ports.BillingClient,
	paymentClient ports.PaymentClient,
) ports.OrchestratorService {
	return &orchestratorService{
		identityClient: identityClient,
		billingClient:  billingClient,
		paymentClient:  paymentClient,
	}
}

func (s *orchestratorService) RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	saga := registration.NewSaga(req, s.identityClient, s.billingClient)
	return saga.Execute(ctx)
}

func (s *orchestratorService) UpgradeSubscription(ctx context.Context, req *dto.UpgradeSubscriptionRequest) (*dto.UpgradeSubscriptionResponse, error) {
	saga := change_plan.NewSaga(req, s.billingClient, s.paymentClient, s.identityClient)
	return saga.Execute(ctx)
}

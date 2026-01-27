package service

import (
	"context"

	"go-link/orchestrator/internal/core/dto"
	"go-link/orchestrator/internal/core/saga/registration"
	"go-link/orchestrator/internal/ports"
)

type orchestratorService struct {
	identityClient ports.IdentityClient
	billingClient  ports.BillingClient
}

func NewOrchestratorService(
	identityClient ports.IdentityClient,
	billingClient ports.BillingClient,
) ports.OrchestratorService {
	return &orchestratorService{
		identityClient: identityClient,
		billingClient:  billingClient,
	}
}

func (s *orchestratorService) RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	saga := registration.NewSaga(req, s.identityClient, s.billingClient)
	return saga.Execute(ctx)
}

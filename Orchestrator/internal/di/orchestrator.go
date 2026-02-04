package di

import (
	driverHttp "go-link/orchestrator/internal/adapters/driver/http"
	"go-link/orchestrator/internal/core/service"
	"go-link/orchestrator/internal/ports"
)

type OrchestratorContainer struct {
	Service ports.OrchestratorService
	Handler driverHttp.OrchestratorHandler
}

func InitOrchestratorDependencies(clientContainer *ClientContainer) *OrchestratorContainer {
	// Service
	svc := service.NewOrchestratorService(
		clientContainer.IdentityClient,
		clientContainer.BillingClient,
		clientContainer.PaymentClient,
	)

	// Handler
	handler := driverHttp.NewOrchestratorHandler(svc)

	return &OrchestratorContainer{
		Service: svc,
		Handler: handler,
	}
}

package di

import (
	"log"

	"go-link/common/pkg/settings"

	grpcDriven "go-link/orchestrator/internal/adapters/driven/grpc"
	"go-link/orchestrator/internal/ports"
)

type ClientContainer struct {
	IdentityClient ports.IdentityClient
	BillingClient  ports.BillingClient
}

func InitClients(cfg settings.Config) *ClientContainer {
	identityClient, err := grpcDriven.NewIdentityClient(cfg.Services.IdentityService)
	if err != nil {
		log.Fatalf("Failed to init Identity Client: %v", err)
	}

	billingClient, err := grpcDriven.NewBillingClient(cfg.Services.BillingService)
	if err != nil {
		log.Fatalf("Failed to init Billing Client: %v", err)
	}

	return &ClientContainer{
		IdentityClient: identityClient,
		BillingClient:  billingClient,
	}
}

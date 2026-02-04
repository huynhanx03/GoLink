package di

import (
	"go.uber.org/zap"

	"go-link/common/pkg/settings"

	"go-link/orchestrator/global"
	grpcDriven "go-link/orchestrator/internal/adapters/driven/grpc"
	"go-link/orchestrator/internal/ports"
)

type ClientContainer struct {
	IdentityClient ports.IdentityClient
	BillingClient  ports.BillingClient
	PaymentClient  ports.PaymentClient
}

func InitClients(cfg settings.Config) *ClientContainer {
	identityClient, err := grpcDriven.NewIdentityClient(cfg.Services.IdentityService)
	if err != nil {
		global.LoggerZap.Fatal("Failed to init Identity Client", zap.Error(err))
	}

	billingClient, err := grpcDriven.NewBillingClient(cfg.Services.BillingService)
	if err != nil {
		global.LoggerZap.Fatal("Failed to init Billing Client", zap.Error(err))
	}

	paymentClient, err := grpcDriven.NewPaymentClient(cfg.Services.PaymentService)
	if err != nil {
		global.LoggerZap.Fatal("Failed to init Payment Client", zap.Error(err))
	}

	return &ClientContainer{
		IdentityClient: identityClient,
		BillingClient:  billingClient,
		PaymentClient:  paymentClient,
	}
}

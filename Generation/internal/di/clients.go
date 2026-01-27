package di

import (
	"go.uber.org/zap"

	billingv1 "go-link/common/gen/go/billing/v1"
	identityv1 "go-link/common/gen/go/identity/v1"
	common_grpc "go-link/common/pkg/grpc"
	"go-link/generation/global"
)

type ClientContainer struct {
	IdentityClient identityv1.IdentityServiceClient
	BillingClient  billingv1.BillingServiceClient
}

func InitClients() *ClientContainer {
	// Identity Client
	identityConn, err := common_grpc.NewClientConn(global.Config.Services.IdentityService)
	if err != nil {
		global.LoggerZap.Fatal("Failed to connect to Identity Service", zap.Error(err))
	}
	identityClient := identityv1.NewIdentityServiceClient(identityConn)

	// Billing Client
	billingConn, err := common_grpc.NewClientConn(global.Config.Services.BillingService)
	if err != nil {
		global.LoggerZap.Fatal("Failed to connect to Billing Service", zap.Error(err))
	}
	billingClient := billingv1.NewBillingServiceClient(billingConn)

	return &ClientContainer{
		IdentityClient: identityClient,
		BillingClient:  billingClient,
	}
}

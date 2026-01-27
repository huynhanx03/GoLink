package grpc

import (
	"go-link/billing/internal/ports"
	billingv1 "go-link/common/gen/go/billing/v1"

	"google.golang.org/grpc"
)

// V1Routes registers the billing service routes
func V1Routes(
	planService ports.PlanService,
) func(srv *grpc.Server) {
	return func(srv *grpc.Server) {
		billingv1.RegisterBillingServiceServer(srv, NewBillingServer(planService))
	}
}

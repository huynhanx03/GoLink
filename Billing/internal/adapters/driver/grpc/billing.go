package grpc

import (
	"context"

	"go-link/billing/internal/constant"
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/ports"
	billingv1 "go-link/common/gen/go/billing/v1"
	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	"go-link/common/pkg/grpc/metadata"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BillingServer struct {
	billingv1.UnimplementedBillingServiceServer
	planService         ports.PlanService
	subscriptionService ports.SubscriptionService
}

func NewBillingServer(planService ports.PlanService, subscriptionService ports.SubscriptionService) *BillingServer {
	return &BillingServer{
		planService:         planService,
		subscriptionService: subscriptionService,
	}
}

func (s *BillingServer) GetTierConfig(ctx context.Context, req *billingv1.GetTierConfigRequest) (*billingv1.GetTierConfigResponse, error) {
	// Extract context (just for consistency, though not strictly used for plan lookup by ID)
	ctx = metadata.ExtractIncomingContext(ctx)

	plan, err := s.planService.Get(ctx, int(req.TierId))
	if err != nil {
		if appErr, ok := err.(*apperr.AppError); ok && appErr.Code == response.CodeNotFound {
			return nil, status.Error(codes.NotFound, "tier not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	maxLinks := int64(-1) // Default to unlimited
	if val, ok := plan.Features[constant.LimitKeyMaxLinks]; ok {
		if limit, ok := val.(float64); ok {
			maxLinks = int64(limit)
		}
	}

	return &billingv1.GetTierConfigResponse{
		TierId:   req.TierId,
		MaxLinks: maxLinks,
	}, nil
}

func (s *BillingServer) CreateSubscription(ctx context.Context, req *billingv1.CreateSubscriptionRequest) (*billingv1.CreateSubscriptionResponse, error) {
	ctx = metadata.ExtractIncomingContext(ctx)

	res, err := s.subscriptionService.Create(ctx, &dto.CreateSubscriptionRequest{
		TenantID: int(req.UserId),
		PlanID:   int(req.PlanId),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &billingv1.CreateSubscriptionResponse{
		SubscriptionId: int64(res.ID),
	}, nil
}

func (s *BillingServer) CancelSubscription(ctx context.Context, req *billingv1.CancelSubscriptionRequest) (*billingv1.CancelSubscriptionResponse, error) {
	ctx = metadata.ExtractIncomingContext(ctx)

	err := s.subscriptionService.Delete(ctx, int(req.SubscriptionId))
	if err != nil {
		return nil, err
	}

	return &billingv1.CancelSubscriptionResponse{
		Success: true,
	}, nil
}

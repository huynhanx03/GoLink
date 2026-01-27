package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/ports"
)

// SubscriptionHandler defines the subscription HTTP handler interface.
type SubscriptionHandler interface {
	Get(ctx context.Context, req *dto.GetSubscriptionRequest) (*dto.SubscriptionResponse, error)
	Create(ctx context.Context, req *dto.CreateSubscriptionRequest) (*dto.SubscriptionResponse, error)
	Update(ctx context.Context, req *dto.UpdateSubscriptionRequest) (*dto.SubscriptionResponse, error)
	Delete(ctx context.Context, req *dto.DeleteSubscriptionRequest) (*dto.SubscriptionResponse, error)
}

type subscriptionHandler struct {
	handler.BaseHandler
	subscriptionService ports.SubscriptionService
}

// NewSubscriptionHandler creates a new SubscriptionHandler instance.
func NewSubscriptionHandler(subscriptionService ports.SubscriptionService) SubscriptionHandler {
	return &subscriptionHandler{
		subscriptionService: subscriptionService,
	}
}

// Get retrieves a subscription by ID.
func (h *subscriptionHandler) Get(ctx context.Context, req *dto.GetSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	return h.subscriptionService.Get(ctx, req.ID)
}

// Create creates a new subscription.
func (h *subscriptionHandler) Create(ctx context.Context, req *dto.CreateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	return h.subscriptionService.Create(ctx, req)
}

// Update updates an existing subscription.
func (h *subscriptionHandler) Update(ctx context.Context, req *dto.UpdateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	return h.subscriptionService.Update(ctx, req.ID, req)
}

// Delete removes a subscription by ID.
func (h *subscriptionHandler) Delete(ctx context.Context, req *dto.DeleteSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	return nil, h.subscriptionService.Delete(ctx, req.ID)
}

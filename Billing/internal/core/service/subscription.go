package service

import (
	"context"
	"strconv"

	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/logger"

	"go-link/billing/internal/constant"
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
	"go-link/billing/internal/core/mapper"
	"go-link/billing/internal/ports"
)

const subscriptionServiceName = "SubscriptionService"

type subscriptionService struct {
	subscriptionRepo ports.SubscriptionRepository
	cache            cache.CacheEngine
	log              *logger.LoggerZap
}

// NewSubscriptionService creates a new SubscriptionService instance.
func NewSubscriptionService(
	subscriptionRepo ports.SubscriptionRepository,
	cache cache.CacheEngine,
	log *logger.LoggerZap,
) ports.SubscriptionService {
	return &subscriptionService{
		subscriptionRepo: subscriptionRepo,
		cache:            cache,
		log:              log,
	}
}

// Get retrieves a subscription by ID.
func (s *subscriptionService) Get(ctx context.Context, id int) (*dto.SubscriptionResponse, error) {
	cacheKey := constant.CacheKeyPrefixSubscriptionID + strconv.Itoa(id)
	var sub *entity.Subscription
	if err := cache.HandleHitCache(ctx, &sub, s.cache, cacheKey); err == nil {
		return mapper.ToSubscriptionResponse(sub), nil
	}

	sub, err := s.subscriptionRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = cache.HandleSetCache(ctx, sub, s.cache, cacheKey, constant.CacheTTLDefault)
	return mapper.ToSubscriptionResponse(sub), nil
}

// Create creates a new subscription.
func (s *subscriptionService) Create(ctx context.Context, req *dto.CreateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	sub := mapper.ToSubscriptionEntityFromCreate(req)

	if err := s.subscriptionRepo.Create(ctx, sub); err != nil {
		return nil, err
	}

	return mapper.ToSubscriptionResponse(sub), nil
}

// Update updates an existing subscription.
func (s *subscriptionService) Update(ctx context.Context, id int, req *dto.UpdateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	sub, err := s.subscriptionRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.PlanID != nil {
		sub.PlanID = *req.PlanID
	}
	if req.Status != nil {
		sub.Status = *req.Status
	}
	if req.CurrentPeriodEnd != nil {
		sub.CurrentPeriodEnd = *req.CurrentPeriodEnd
	}
	if req.CancelAtPeriodEnd != nil {
		sub.CancelAtPeriodEnd = *req.CancelAtPeriodEnd
	}

	if err := s.subscriptionRepo.Update(ctx, sub); err != nil {
		return nil, err
	}

	cache.HandleUpdateCache(ctx, sub, s.cache, constant.CacheKeyPrefixSubscriptionID+strconv.Itoa(id), constant.CacheTTLDefault)

	return mapper.ToSubscriptionResponse(sub), nil
}

// Delete removes a subscription by ID.
func (s *subscriptionService) Delete(ctx context.Context, id int) error {
	if err := s.subscriptionRepo.Delete(ctx, id); err != nil {
		return err
	}

	_ = cache.HandleDeleteCache(ctx, s.cache, constant.CacheKeyPrefixSubscriptionID+strconv.Itoa(id))

	return nil
}

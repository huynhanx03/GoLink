package service

import (
	"context"
	"strconv"
	"time"

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
	planRepo         ports.PlanRepository
	cache            cache.CacheEngine
	log              *logger.LoggerZap
}

// NewSubscriptionService creates a new SubscriptionService instance.
func NewSubscriptionService(
	subscriptionRepo ports.SubscriptionRepository,
	planRepo ports.PlanRepository,
	cache cache.CacheEngine,
	log *logger.LoggerZap,
) ports.SubscriptionService {
	return &subscriptionService{
		subscriptionRepo: subscriptionRepo,
		planRepo:         planRepo,
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
	plan, err := s.planRepo.Get(ctx, req.PlanID)
	if err != nil {
		return nil, err
	}

	now, periodEnd := s.calculatePeriod(plan)

	sub := &entity.Subscription{
		TenantID:           req.TenantID,
		PlanID:             req.PlanID,
		Status:             constant.SubscriptionStatusActive,
		CurrentPeriodStart: now,
		CurrentPeriodEnd:   periodEnd,
		CancelAtPeriodEnd:  true,
	}

	if err := s.subscriptionRepo.Create(ctx, sub); err != nil {
		return nil, err
	}

	return mapper.ToSubscriptionResponse(sub), nil
}

// UpdateByTenant updates a subscription by tenant ID.
func (s *subscriptionService) UpdateByTenant(ctx context.Context, tenantID int, req *dto.UpdateSubscriptionRequest) (*dto.SubscriptionResponse, error) {
	sub, err := s.subscriptionRepo.GetByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	if req.PlanID != nil {
		sub.PlanID = *req.PlanID

		plan, err := s.planRepo.Get(ctx, *req.PlanID)
		if err != nil {
			return nil, err
		}

		now, periodEnd := s.calculatePeriod(plan)
		sub.CurrentPeriodStart = now
		sub.CurrentPeriodEnd = periodEnd
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

	cache.HandleUpdateCache(ctx, sub, s.cache, constant.CacheKeyPrefixSubscriptionID+strconv.Itoa(sub.ID), constant.CacheTTLDefault)

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

func (s *subscriptionService) calculatePeriod(plan *entity.Plan) (time.Time, time.Time) {
	now := time.Now()
	var periodEnd time.Time

	switch plan.Period {
	case constant.PlanPeriodForever:
		periodEnd = now.AddDate(100, 0, 0)
	case constant.PlanPeriodYear:
		periodEnd = now.AddDate(1, 0, 0)
	case constant.PlanPeriodMonth:
		fallthrough
	default:
		periodEnd = now.AddDate(0, 1, 0)
	}
	return now, periodEnd
}

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

const planServiceName = "PlanService"

type planService struct {
	planRepo ports.PlanRepository
	cache    cache.LocalCache[string, any]
	log      *logger.LoggerZap
}

// NewPlanService creates a new PlanService instance.
func NewPlanService(
	planRepo ports.PlanRepository,
	cache cache.LocalCache[string, any],
	log *logger.LoggerZap,
) ports.PlanService {
	return &planService{
		planRepo: planRepo,
		cache:    cache,
		log:      log,
	}
}

// Get retrieves a plan by ID.
func (s *planService) Get(ctx context.Context, id int) (*dto.PlanResponse, error) {
	cacheKey := constant.CacheKeyPrefixPlanID + strconv.Itoa(id)
	if v, found := cache.GetLocal[*entity.Plan](s.cache, cacheKey); found {
		return mapper.ToPlanResponse(v), nil
	}

	plan, err := s.planRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	cache.SetLocal(s.cache, cacheKey, plan, constant.CacheCostID)
	return mapper.ToPlanResponse(plan), nil
}

// Create creates a new plan.
func (s *planService) Create(ctx context.Context, req *dto.CreatePlanRequest) (*dto.PlanResponse, error) {
	plan := mapper.ToPlanEntityFromCreate(req)

	if err := s.planRepo.Create(ctx, plan); err != nil {
		return nil, err
	}

	return mapper.ToPlanResponse(plan), nil
}

// Update updates an existing plan.
func (s *planService) Update(ctx context.Context, id int, req *dto.UpdatePlanRequest) (*dto.PlanResponse, error) {
	plan, err := s.planRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		plan.Name = *req.Name
	}
	if req.BasePrice != nil {
		plan.BasePrice = *req.BasePrice
	}
	if req.Limits != nil {
		plan.Limits = *req.Limits
	}
	if req.IsActive != nil {
		plan.IsActive = *req.IsActive
	}

	if err := s.planRepo.Update(ctx, plan); err != nil {
		return nil, err
	}

	cache.UpdateLocal(s.cache, constant.CacheKeyPrefixPlanID+strconv.Itoa(id), plan, constant.CacheCostID)

	return mapper.ToPlanResponse(plan), nil
}

// Delete removes a plan by ID.
func (s *planService) Delete(ctx context.Context, id int) error {
	if err := s.planRepo.Delete(ctx, id); err != nil {
		return err
	}

	cache.DeleteLocal(s.cache, constant.CacheKeyPrefixPlanID+strconv.Itoa(id))

	return nil
}

// FindAll retrieves all plans.
func (s *planService) FindAll(ctx context.Context) ([]*dto.PlanResponse, error) {
	plans, err := s.planRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.PlanResponse, 0, len(plans))
	for _, p := range plans {
		responses = append(responses, mapper.ToPlanResponse(p))
	}

	return responses, nil
}

// FindActive retrieves all active plans.
func (s *planService) FindActive(ctx context.Context) ([]*dto.PlanResponse, error) {
	plans, err := s.planRepo.FindActive(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.PlanResponse, 0, len(plans))
	for _, p := range plans {
		responses = append(responses, mapper.ToPlanResponse(p))
	}

	return responses, nil
}

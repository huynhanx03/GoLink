package service

import (
	"context"
	"net/http"
	"strconv"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/common/http/response"

	"go-link/identity/internal/constant"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/core/mapper"
	"go-link/identity/internal/ports"
)

const tenantServiceName = "TenantService"

type tenantService struct {
	tenantRepo       ports.TenantRepository
	tenantMemberRepo ports.TenantMemberRepository
	cache            cache.LocalCache[string, any]
}

// NewTenantService creates a new TenantService instance.
func NewTenantService(
	tenantRepo ports.TenantRepository,
	tenantMemberRepo ports.TenantMemberRepository,
	cache cache.LocalCache[string, any],
) ports.TenantService {
	return &tenantService{
		tenantRepo:       tenantRepo,
		tenantMemberRepo: tenantMemberRepo,
		cache:            cache,
	}
}

// Get retrieves a tenant by ID.
func (s *tenantService) Get(ctx context.Context, id int) (*dto.TenantResponse, error) {
	cacheKey := constant.CacheKeyPrefixTenantID + strconv.Itoa(id)
	if t, found := cache.GetLocal[*entity.Tenant](s.cache, cacheKey); found {
		return mapper.ToTenantResponse(t), nil
	}

	tenant, err := s.tenantRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	cache.SetLocal(s.cache, cacheKey, tenant, constant.CacheCostID)
	return mapper.ToTenantResponse(tenant), nil
}

// Create creates a new tenant.
func (s *tenantService) Create(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	tenant := mapper.ToTenantEntityFromCreate(req)

	if err := s.tenantRepo.Create(ctx, tenant); err != nil {
		return nil, err
	}

	return mapper.ToTenantResponse(tenant), nil
}

// Update updates an existing tenant.
func (s *tenantService) Update(ctx context.Context, id int, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error) {
	tenant, err := s.tenantRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		tenant.Name = *req.Name
	}
	if req.PlanID != nil {
		tenant.PlanID = *req.PlanID
	}

	tenant.ID = id
	if err := s.tenantRepo.Update(ctx, tenant); err != nil {
		return nil, err
	}

	cacheKeyID := constant.CacheKeyPrefixTenantID + strconv.Itoa(id)
	cache.SetLocal(s.cache, cacheKeyID, tenant, constant.CacheCostID)

	return mapper.ToTenantResponse(tenant), nil
}

// Delete removes a tenant by ID.
func (s *tenantService) Delete(ctx context.Context, id int) error {
	exists, err := s.tenantRepo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return apperr.NewError(tenantServiceName, response.CodeNotFound, apperr.MsgNotFound, http.StatusNotFound, nil)
	}

	if err := s.tenantRepo.Delete(ctx, id); err != nil {
		return err
	}

	// Invalidate Cache
	cacheKeyID := constant.CacheKeyPrefixTenantID + strconv.Itoa(id)
	cache.DeleteLocal(s.cache, cacheKeyID)

	return nil
}

// GetByUserID retrieves all tenants that a user belongs to.
func (s *tenantService) GetByUserID(ctx context.Context, userID int) ([]*dto.TenantResponse, error) {
	memberships, err := s.tenantMemberRepo.GetByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if len(memberships) == 0 {
		return []*dto.TenantResponse{}, nil
	}

	tenantIDs := make([]int, len(memberships))
	for i, m := range memberships {
		tenantIDs[i] = m.TenantID
	}

	tenants, err := s.tenantRepo.GetByIDs(ctx, tenantIDs)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.TenantResponse, len(tenants))
	for i, t := range tenants {
		responses[i] = mapper.ToTenantResponse(t)
	}

	return responses, nil
}

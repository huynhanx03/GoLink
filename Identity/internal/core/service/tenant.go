package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/mapper"
	"go-link/identity/internal/ports"
)

type tenantService struct {
	tenantRepo ports.TenantRepository
}

// NewTenantService creates a new TenantService instance.
func NewTenantService(tenantRepo ports.TenantRepository) ports.TenantService {
	return &tenantService{tenantRepo: tenantRepo}
}

// Get retrieves a tenant by ID.
func (s *tenantService) Get(ctx context.Context, id int) (*dto.TenantResponse, error) {
	tenant, err := s.tenantRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get tenant", http.StatusInternalServerError)
	}
	return mapper.ToTenantResponse(tenant), nil
}

// Create creates a new tenant.
func (s *tenantService) Create(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	tenant := mapper.ToTenantEntityFromCreate(req)

	if err := s.tenantRepo.Create(ctx, tenant); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create tenant", http.StatusInternalServerError)
	}

	return mapper.ToTenantResponse(tenant), nil
}

// Update updates an existing tenant.
func (s *tenantService) Update(ctx context.Context, id int, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error) {
	tenant, err := s.tenantRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get tenant", http.StatusInternalServerError)
	}

	if req.Name != nil {
		tenant.Name = *req.Name
	}
	if req.TierID != nil {
		tenant.TierID = *req.TierID
	}

	tenant.ID = id
	if err := s.tenantRepo.Update(ctx, tenant); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to update tenant", http.StatusInternalServerError)
	}

	return mapper.ToTenantResponse(tenant), nil
}

// Delete removes a tenant by ID.
func (s *tenantService) Delete(ctx context.Context, id int) error {
	exists, err := s.tenantRepo.Exists(ctx, id)
	if err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to check tenant exists", http.StatusInternalServerError)
	}

	if !exists {
		return apperr.New(response.CodeNotFound, "tenant not found", http.StatusNotFound, nil)
	}

	if err := s.tenantRepo.Delete(ctx, id); err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to delete tenant", http.StatusInternalServerError)
	}

	return nil
}

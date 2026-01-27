package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/ports"
)

// TenantHandler defines the tenant HTTP handler interface.
type TenantHandler interface {
	Get(ctx context.Context, req *dto.GetTenantRequest) (*dto.TenantResponse, error)
	Create(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
	Update(ctx context.Context, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error)
	Delete(ctx context.Context, req *dto.DeleteTenantRequest) (*dto.TenantResponse, error)
}

type tenantHandler struct {
	handler.BaseHandler
	tenantService ports.TenantService
}

// NewTenantHandler creates a new TenantHandler instance.
func NewTenantHandler(tenantService ports.TenantService) TenantHandler {
	return &tenantHandler{
		tenantService: tenantService,
	}
}

// Get retrieves a tenant by ID.
func (h *tenantHandler) Get(ctx context.Context, req *dto.GetTenantRequest) (*dto.TenantResponse, error) {
	return h.tenantService.Get(ctx, req.ID)
}

// Create creates a new tenant.
func (h *tenantHandler) Create(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	return h.tenantService.Create(ctx, req)
}

// Update updates an existing tenant.
func (h *tenantHandler) Update(ctx context.Context, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error) {
	return h.tenantService.Update(ctx, req.ID, req)
}

// Delete removes a tenant by ID.
func (h *tenantHandler) Delete(ctx context.Context, req *dto.DeleteTenantRequest) (*dto.TenantResponse, error) {
	return nil, h.tenantService.Delete(ctx, req.ID)
}

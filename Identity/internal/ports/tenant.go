package ports

import (
	"context"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
)

// TenantRepository defines the tenant data access interface.
type TenantRepository interface {
	Get(ctx context.Context, id int) (*entity.Tenant, error)
	Create(ctx context.Context, e *entity.Tenant) error
	Update(ctx context.Context, e *entity.Tenant) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) (bool, error)
}

// TenantService defines the tenant business logic interface.
type TenantService interface {
	Get(ctx context.Context, id int) (*dto.TenantResponse, error)
	Create(ctx context.Context, req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdateTenantRequest) (*dto.TenantResponse, error)
	Delete(ctx context.Context, id int) error
}

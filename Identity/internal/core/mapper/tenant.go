package mapper

import (
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
)

// ToTenantResponse converts Tenant entity to TenantResponse DTO.
func ToTenantResponse(e *entity.Tenant) *dto.TenantResponse {
	if e == nil {
		return nil
	}
	return &dto.TenantResponse{
		ID:     e.ID,
		Name:   e.Name,
		TierID: e.TierID,
	}
}

// ToTenantEntityFromCreate converts CreateTenantRequest to Tenant entity.
func ToTenantEntityFromCreate(req *dto.CreateTenantRequest) *entity.Tenant {
	return &entity.Tenant{
		Name:   req.Name,
		TierID: 0,
	}
}

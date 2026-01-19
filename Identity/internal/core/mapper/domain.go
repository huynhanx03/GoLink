package mapper

import (
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
)

// ToDomainResponse converts Domain entity to DomainResponse DTO.
func ToDomainResponse(e *entity.Domain) *dto.DomainResponse {
	if e == nil {
		return nil
	}
	return &dto.DomainResponse{
		ID:         e.ID,
		Domain:     e.Domain,
		IsVerified: e.IsVerified,
		TenantID:   e.TenantID,
	}
}

// ToDomainEntityFromCreate converts CreateDomainRequest to Domain entity.
func ToDomainEntityFromCreate(req *dto.CreateDomainRequest) *entity.Domain {
	return &entity.Domain{
		Domain:   req.Domain,
		TenantID: req.TenantID,
	}
}

package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/ports"
)

// DomainHandler defines the domain HTTP handler interface.
type DomainHandler interface {
	Find(ctx context.Context, req *d.QueryOptions) (*d.Paginated[*dto.DomainResponse], error)
	Get(ctx context.Context, req *dto.GetDomainRequest) (*dto.DomainResponse, error)
	Create(ctx context.Context, req *dto.CreateDomainRequest) (*dto.DomainResponse, error)
	Update(ctx context.Context, req *dto.UpdateDomainRequest) (*dto.DomainResponse, error)
	Delete(ctx context.Context, req *dto.DeleteDomainRequest) (*dto.DomainResponse, error)
}

type domainHandler struct {
	handler.BaseHandler
	domainService ports.DomainService
}

// NewDomainHandler creates a new DomainHandler instance.
func NewDomainHandler(domainService ports.DomainService) DomainHandler {
	return &domainHandler{
		domainService: domainService,
	}
}

// Find retrieves domains with pagination.
func (h *domainHandler) Find(ctx context.Context, req *d.QueryOptions) (*d.Paginated[*dto.DomainResponse], error) {
	return h.domainService.Find(ctx, req)
}

// Get retrieves a domain by ID.
func (h *domainHandler) Get(ctx context.Context, req *dto.GetDomainRequest) (*dto.DomainResponse, error) {
	return h.domainService.Get(ctx, req.ID)
}

// Create creates a new domain.
func (h *domainHandler) Create(ctx context.Context, req *dto.CreateDomainRequest) (*dto.DomainResponse, error) {
	return h.domainService.Create(ctx, req)
}

// Update updates an existing domain.
func (h *domainHandler) Update(ctx context.Context, req *dto.UpdateDomainRequest) (*dto.DomainResponse, error) {
	return h.domainService.Update(ctx, req.ID, req)
}

// Delete removes a domain by ID.
func (h *domainHandler) Delete(ctx context.Context, req *dto.DeleteDomainRequest) (*dto.DomainResponse, error) {
	return nil, h.domainService.Delete(ctx, req.ID)
}

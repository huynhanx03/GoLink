package service

import (
	"context"
	"net/http"
	"strconv"

	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/common/http/response"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/constant"
	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/core/mapper"
	"go-link/identity/internal/ports"
)

const domainServiceName = "DomainService"

type domainService struct {
	domainRepo ports.DomainRepository
	cache      cache.LocalCache[string, any]
}

// NewDomainService creates a new DomainService instance.
func NewDomainService(domainRepo ports.DomainRepository, cache cache.LocalCache[string, any]) ports.DomainService {
	return &domainService{domainRepo: domainRepo, cache: cache}
}

// Find retrieves domains with pagination.
func (s *domainService) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*dto.DomainResponse], error) {
	domains, err := s.domainRepo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	if domains.Records == nil {
		return &d.Paginated[*dto.DomainResponse]{
			Records:    &[]*dto.DomainResponse{},
			Pagination: domains.Pagination,
		}, nil
	}

	entities := *domains.Records
	responses := make([]*dto.DomainResponse, len(entities))
	for i, domain := range entities {
		responses[i] = mapper.ToDomainResponse(domain)
	}

	return &d.Paginated[*dto.DomainResponse]{
		Records:    &responses,
		Pagination: domains.Pagination,
	}, nil
}

// Get retrieves a domain by ID.
func (s *domainService) Get(ctx context.Context, id int) (*dto.DomainResponse, error) {
	cacheKey := constant.CacheKeyPrefixDomainID + strconv.Itoa(id)
	if d, found := cache.GetLocal[*entity.Domain](s.cache, cacheKey); found {
		return mapper.ToDomainResponse(d), nil
	}

	domain, err := s.domainRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	cache.SetLocal(s.cache, cacheKey, domain, constant.CacheCostID)
	return mapper.ToDomainResponse(domain), nil
}

// Create creates a new domain.
func (s *domainService) Create(ctx context.Context, req *dto.CreateDomainRequest) (*dto.DomainResponse, error) {
	domain := mapper.ToDomainEntityFromCreate(req)
	domain.IsVerified = false
	if err := s.domainRepo.Create(ctx, domain); err != nil {
		return nil, err
	}

	return mapper.ToDomainResponse(domain), nil
}

// Update updates an existing domain.
func (s *domainService) Update(ctx context.Context, id int, req *dto.UpdateDomainRequest) (*dto.DomainResponse, error) {
	domain, err := s.domainRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Domain != nil {
		domain.Domain = *req.Domain
	}
	if req.IsVerified != nil {
		domain.IsVerified = *req.IsVerified
	}

	domain.ID = id
	if err := s.domainRepo.Update(ctx, domain); err != nil {
		return nil, err
	}

	// Invalidate Cache
	cacheKeyID := constant.CacheKeyPrefixDomainID + strconv.Itoa(id)
	cache.SetLocal(s.cache, cacheKeyID, domain, constant.CacheCostID)

	return mapper.ToDomainResponse(domain), nil
}

// Delete removes a domain by ID.
func (s *domainService) Delete(ctx context.Context, id int) error {
	exists, err := s.domainRepo.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		return NewError(domainServiceName, response.CodeNotFound, "domain not found", http.StatusNotFound, nil)
	}

	if err := s.domainRepo.Delete(ctx, id); err != nil {
		return err
	}

	if err := s.domainRepo.Delete(ctx, id); err != nil {
		return err
	}

	// Invalidate Cache
	cacheKeyID := constant.CacheKeyPrefixDomainID + strconv.Itoa(id)
	cache.DeleteLocal(s.cache, cacheKeyID)

	return nil
}

package service

import (
	"context"
	"net/http"

	"go-link/common/pkg/common/apperr"
	"go-link/common/pkg/common/http/response"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/mapper"
	"go-link/identity/internal/ports"
)

type resourceService struct {
	resourceRepo ports.ResourceRepository
	cacheService ports.CacheService
}

// NewResourceService creates a new ResourceService instance.
func NewResourceService(resourceRepo ports.ResourceRepository, cacheService ports.CacheService) ports.ResourceService {
	return &resourceService{resourceRepo: resourceRepo, cacheService: cacheService}
}

// Find retrieves resources with pagination.
func (s *resourceService) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*dto.ResourceResponse], error) {
	resources, err := s.resourceRepo.Find(ctx, opts)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to find resources", http.StatusInternalServerError)
	}

	if resources.Records == nil {
		return &d.Paginated[*dto.ResourceResponse]{
			Records:    &[]*dto.ResourceResponse{},
			Pagination: resources.Pagination,
		}, nil
	}

	entities := *resources.Records
	responses := make([]*dto.ResourceResponse, len(entities))
	for i, resource := range entities {
		responses[i] = mapper.ToResourceResponse(resource)
	}

	return &d.Paginated[*dto.ResourceResponse]{
		Records:    &responses,
		Pagination: resources.Pagination,
	}, nil
}

// Get retrieves a resource by ID.
func (s *resourceService) Get(ctx context.Context, id int) (*dto.ResourceResponse, error) {
	resource, err := s.resourceRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get resource", http.StatusInternalServerError)
	}

	return mapper.ToResourceResponse(resource), nil
}

// Create creates a new resource.
func (s *resourceService) Create(ctx context.Context, req *dto.CreateResourceRequest) (*dto.ResourceResponse, error) {
	resource := mapper.ToResourceEntityFromCreate(req)

	if err := s.resourceRepo.Create(ctx, resource); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create resource", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return mapper.ToResourceResponse(resource), nil
}

// Update updates an existing resource.
func (s *resourceService) Update(ctx context.Context, id int, req *dto.UpdateResourceRequest) (*dto.ResourceResponse, error) {
	resource, err := s.resourceRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get resource", http.StatusInternalServerError)
	}

	if req.Key != nil {
		resource.Key = *req.Key
	}
	if req.Description != nil {
		resource.Description = req.Description
	}

	resource.ID = id
	if err := s.resourceRepo.Update(ctx, resource); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to update resource", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return mapper.ToResourceResponse(resource), nil
}

// Delete removes a resource by ID.
func (s *resourceService) Delete(ctx context.Context, id int) error {
	exists, err := s.resourceRepo.Exists(ctx, id)
	if err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to check resource exists", http.StatusInternalServerError)
	}

	if !exists {
		return apperr.New(response.CodeNotFound, "resource not found", http.StatusNotFound, nil)
	}

	if err := s.resourceRepo.Delete(ctx, id); err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to delete resource", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return nil
}

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

type permissionService struct {
	permissionRepo ports.PermissionRepository
	cacheService   ports.CacheService
}

// NewPermissionService creates a new PermissionService instance.
func NewPermissionService(permissionRepo ports.PermissionRepository, cacheService ports.CacheService) ports.PermissionService {
	return &permissionService{permissionRepo: permissionRepo, cacheService: cacheService}
}

// Find retrieves permissions with pagination.
func (s *permissionService) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*dto.PermissionResponse], error) {
	permissions, err := s.permissionRepo.Find(ctx, opts)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to find permissions", http.StatusInternalServerError)
	}

	if permissions.Records == nil {
		return &d.Paginated[*dto.PermissionResponse]{
			Records:    &[]*dto.PermissionResponse{},
			Pagination: permissions.Pagination,
		}, nil
	}

	entities := *permissions.Records
	responses := make([]*dto.PermissionResponse, len(entities))
	for i, permission := range entities {
		responses[i] = mapper.ToPermissionResponse(permission)
	}

	return &d.Paginated[*dto.PermissionResponse]{
		Records:    &responses,
		Pagination: permissions.Pagination,
	}, nil
}

// Get retrieves a permission by ID
func (s *permissionService) Get(ctx context.Context, id int) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get permission", http.StatusInternalServerError)
	}

	return mapper.ToPermissionResponse(permission), nil
}

// Create creates a new permission.
func (s *permissionService) Create(ctx context.Context, req *dto.CreatePermissionRequest) (*dto.PermissionResponse, error) {
	permission := mapper.ToPermissionEntityFromCreate(req)

	if err := s.permissionRepo.Create(ctx, permission); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create permission", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return mapper.ToPermissionResponse(permission), nil
}

// Update updates an existing permission.
func (s *permissionService) Update(ctx context.Context, id int, req *dto.UpdatePermissionRequest) (*dto.PermissionResponse, error) {
	permission, err := s.permissionRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get permission", http.StatusInternalServerError)
	}

	if req.Description != nil {
		permission.Description = req.Description
	}
	if req.Scopes != nil {
		permission.Scopes = *req.Scopes
	}

	permission.ID = id
	if err := s.permissionRepo.Update(ctx, permission); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to update permission", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return mapper.ToPermissionResponse(permission), nil
}

// Delete removes a permission by ID.
func (s *permissionService) Delete(ctx context.Context, id int) error {
	exists, err := s.permissionRepo.Exists(ctx, id)
	if err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to check permission exists", http.StatusInternalServerError)
	}

	if !exists {
		return apperr.New(response.CodeNotFound, "permission not found", http.StatusNotFound, nil)
	}

	if err := s.permissionRepo.Delete(ctx, id); err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to delete permission", http.StatusInternalServerError)
	}

	// Invalidate Permission Config Version
	if err := s.cacheService.InvalidatePermissionConfig(ctx); err != nil {
		// Log error but don't fail request
	}

	return nil
}

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

type attributeDefinitionService struct {
	attrDefRepo ports.AttributeDefinitionRepository
}

// NewAttributeDefinitionService creates a new AttributeDefinitionService instance.
func NewAttributeDefinitionService(attrDefRepo ports.AttributeDefinitionRepository) ports.AttributeDefinitionService {
	return &attributeDefinitionService{attrDefRepo: attrDefRepo}
}

// Find retrieves attribute definitions with pagination.
func (s *attributeDefinitionService) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*dto.AttributeDefinitionResponse], error) {
	result, err := s.attrDefRepo.Find(ctx, opts)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to find attribute definitions", http.StatusInternalServerError)
	}

	if result.Records == nil {
		return &d.Paginated[*dto.AttributeDefinitionResponse]{
			Records:    &[]*dto.AttributeDefinitionResponse{},
			Pagination: result.Pagination,
		}, nil
	}

	entities := *result.Records
	responses := make([]*dto.AttributeDefinitionResponse, len(entities))
	for i, e := range entities {
		responses[i] = mapper.ToAttributeDefinitionResponse(e)
	}

	return &d.Paginated[*dto.AttributeDefinitionResponse]{
		Records:    &responses,
		Pagination: result.Pagination,
	}, nil
}

// Get retrieves an attribute definition by ID.
func (s *attributeDefinitionService) Get(ctx context.Context, id int) (*dto.AttributeDefinitionResponse, error) {
	attrDef, err := s.attrDefRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get attribute definition", http.StatusInternalServerError)
	}
	return mapper.ToAttributeDefinitionResponse(attrDef), nil
}

// Create creates a new attribute definition.
func (s *attributeDefinitionService) Create(ctx context.Context, req *dto.CreateAttributeDefinitionRequest) (*dto.AttributeDefinitionResponse, error) {
	attrDef := mapper.ToAttributeDefinitionEntityFromCreate(req)
	if err := s.attrDefRepo.Create(ctx, attrDef); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to create attribute definition", http.StatusInternalServerError)
	}
	return mapper.ToAttributeDefinitionResponse(attrDef), nil
}

// Update updates an existing attribute definition.
func (s *attributeDefinitionService) Update(ctx context.Context, id int, req *dto.UpdateAttributeDefinitionRequest) (*dto.AttributeDefinitionResponse, error) {
	attrDef, err := s.attrDefRepo.Get(ctx, id)
	if err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to get attribute definition", http.StatusInternalServerError)
	}

	if req.Key != nil {
		attrDef.Key = *req.Key
	}
	if req.DataType != nil {
		attrDef.DataType = *req.DataType
	}
	if req.Description != nil {
		attrDef.Description = *req.Description
	}

	attrDef.ID = id
	if err := s.attrDefRepo.Update(ctx, attrDef); err != nil {
		return nil, apperr.Wrap(err, response.CodeDatabaseError, "failed to update attribute definition", http.StatusInternalServerError)
	}

	return mapper.ToAttributeDefinitionResponse(attrDef), nil
}

// Delete removes an attribute definition by ID.
func (s *attributeDefinitionService) Delete(ctx context.Context, id int) error {
	exists, err := s.attrDefRepo.Exists(ctx, id)
	if err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to check attribute definition exists", http.StatusInternalServerError)
	}

	if !exists {
		return apperr.New(response.CodeNotFound, "attribute definition not found", http.StatusNotFound, nil)
	}

	if err := s.attrDefRepo.Delete(ctx, id); err != nil {
		return apperr.Wrap(err, response.CodeDatabaseError, "failed to delete attribute definition", http.StatusInternalServerError)
	}

	return nil
}

package db

import (
	"context"

	"go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driven/db/ent/generate/attributedefinition"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

type AttributeDefinitionRepository struct {
	repo   *ent.BaseRepository[generate.AttributeDefinition, *generate.AttributeDefinition, int]
	client *generate.AttributeDefinitionClient
}

// NewAttributeDefinitionRepository creates a new AttributeDefinitionRepository instance.
func NewAttributeDefinitionRepository(client interface{}) ports.AttributeDefinitionRepository {
	entClient := client.(*generate.Client)
	return &AttributeDefinitionRepository{
		repo:   ent.NewBaseRepository[generate.AttributeDefinition, *generate.AttributeDefinition, int](client),
		client: entClient.AttributeDefinition,
	}
}

// Get retrieves an attribute definition by ID.
func (r *AttributeDefinitionRepository) Get(ctx context.Context, id int) (*entity.AttributeDefinition, error) {
	record, err := r.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.ToAttributeDefinitionEntity(record), nil
}

// GetByKey retrieves an attribute definition by key.
func (r *AttributeDefinitionRepository) GetByKey(ctx context.Context, key string) (*entity.AttributeDefinition, error) {
	record, err := r.client.Query().
		Where(attributedefinition.Key(key)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.ToAttributeDefinitionEntity(record), nil
}

// FindAll retrieves all attribute definitions.
func (r *AttributeDefinitionRepository) FindAll(ctx context.Context) ([]*entity.AttributeDefinition, error) {
	records, err := r.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	entities := make([]*entity.AttributeDefinition, len(records))
	for i, record := range records {
		entities[i] = mapper.ToAttributeDefinitionEntity(record)
	}
	return entities, nil
}

// Find retrieves attribute definitions with pagination.
func (r *AttributeDefinitionRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.AttributeDefinition], error) {
	result, err := r.repo.Find(ctx, opts)
	if err != nil {
		return nil, err
	}
	entities := make([]*entity.AttributeDefinition, len(*result.Records))
	for i, record := range *result.Records {
		entities[i] = mapper.ToAttributeDefinitionEntity(record)
	}
	return &d.Paginated[*entity.AttributeDefinition]{
		Records:    &entities,
		Pagination: result.Pagination,
	}, nil
}

// Create creates a new attribute definition.
func (r *AttributeDefinitionRepository) Create(ctx context.Context, e *entity.AttributeDefinition) error {
	model := mapper.ToAttributeDefinitionModel(e)
	if err := r.repo.Create(ctx, model); err != nil {
		return err
	}

	if created := mapper.ToAttributeDefinitionEntity(model); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing attribute definition.
func (r *AttributeDefinitionRepository) Update(ctx context.Context, e *entity.AttributeDefinition) error {
	model := mapper.ToAttributeDefinitionModel(e)
	if err := r.repo.Update(ctx, model); err != nil {
		return err
	}
	e.UpdatedAt = model.UpdatedAt
	return nil
}

// Delete removes an attribute definition by ID.
func (r *AttributeDefinitionRepository) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

// Exists checks if an attribute definition exists by ID.
func (r *AttributeDefinitionRepository) Exists(ctx context.Context, id int) (bool, error) {
	return r.repo.Exists(ctx, id)
}

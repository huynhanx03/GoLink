package db

import (
	"context"

	commonEnt "go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"

	"entgo.io/ent/dialect/sql"

	"go-link/identity/internal/adapters/driven/db/ent/builder"
	"go-link/identity/internal/adapters/driven/db/ent/generate/domain"

	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

const domainRepoName = "DomainRepository"

type DomainRepository struct {
	client *dbEnt.EntClient
}

// NewDomainRepository creates a new DomainRepository instance.
func NewDomainRepository(client *dbEnt.EntClient) ports.DomainRepository {
	return &DomainRepository{
		client: client,
	}
}

// Find retrieves domains with pagination.
func (r *DomainRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Domain], error) {
	client := r.client.DB(ctx)

	query := client.Domain.Query()
	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplyFilters(opts.Filters, s)
		})
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, domainRepoName)
	}

	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplySort(opts.Sort, s)
			commonEnt.ApplyPagination(opts.Pagination, s)
		})
	}

	records, err := query.All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, domainRepoName)
	}

	entities := make([]*entity.Domain, len(records))
	for i, record := range records {
		entities[i] = mapper.ToDomainEntity(record)
	}

	paginationOpts := &d.PaginationOptions{}
	if opts != nil && opts.Pagination != nil {
		paginationOpts = opts.Pagination
	} else {
		paginationOpts.SetDefaults()
	}

	meta := d.CalculatePagination(
		paginationOpts.Page,
		paginationOpts.PageSize,
		int64(total),
	)

	return &d.Paginated[*entity.Domain]{
		Records:    &entities,
		Pagination: meta,
	}, nil
}

// Get retrieves a domain by ID.
func (r *DomainRepository) Get(ctx context.Context, id int) (*entity.Domain, error) {
	record, err := r.client.DB(ctx).Domain.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, domainRepoName)
	}
	return mapper.ToDomainEntity(record), nil
}

// Create creates a new domain.
func (r *DomainRepository) Create(ctx context.Context, e *entity.Domain) error {
	create := builder.BuildCreateDomain(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, domainRepoName)
	}

	if created := mapper.ToDomainEntity(record); created != nil {
		*e = *created
	}
	return nil
}

// Update updates an existing domain.
func (r *DomainRepository) Update(ctx context.Context, e *entity.Domain) error {
	update := builder.BuildUpdateDomain(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, domainRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

// Delete removes a domain by ID.
func (r *DomainRepository) Delete(ctx context.Context, id int) error {
	if err := r.client.DB(ctx).Domain.DeleteOneID(id).Exec(ctx); err != nil {
		return commonEnt.MapEntError(err, domainRepoName)
	}
	return nil
}

// Exists checks if a domain exists by ID.
func (r *DomainRepository) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := r.client.DB(ctx).Domain.Query().Where(domain.ID(id)).Exist(ctx)
	return exists, commonEnt.MapEntError(err, domainRepoName)
}

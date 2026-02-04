package db

import (
	"context"

	commonEnt "go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"

	"entgo.io/ent/dialect/sql"

	"go-link/identity/internal/adapters/driven/db/ent/builder"
	"go-link/identity/internal/adapters/driven/db/ent/generate/tenant"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

const tenantRepoName = "TenantRepository"

type TenantRepository struct {
	client *dbEnt.EntClient
}

func NewTenantRepository(client *dbEnt.EntClient) ports.TenantRepository {
	return &TenantRepository{client: client}
}

func (r *TenantRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Tenant], error) {
	client := r.client.DB(ctx)

	query := client.Tenant.Query()
	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplyFilters(opts.Filters, s)
		})
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantRepoName)
	}

	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplySort(opts.Sort, s)
			commonEnt.ApplyPagination(opts.Pagination, s)
		})
	}

	records, err := query.All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantRepoName)
	}

	entities := make([]*entity.Tenant, len(records))
	for i, record := range records {
		entities[i] = mapper.ToTenantEntity(record)
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

	return &d.Paginated[*entity.Tenant]{
		Records:    &entities,
		Pagination: meta,
	}, nil
}

func (r *TenantRepository) Get(ctx context.Context, id int) (*entity.Tenant, error) {
	record, err := r.client.DB(ctx).Tenant.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantRepoName)
	}
	return mapper.ToTenantEntity(record), nil
}

func (r *TenantRepository) GetByIDs(ctx context.Context, ids []int) ([]*entity.Tenant, error) {
	records, err := r.client.DB(ctx).Tenant.Query().
		Where(tenant.IDIn(ids...)).
		All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantRepoName)
	}

	entities := make([]*entity.Tenant, len(records))
	for i, record := range records {
		entities[i] = mapper.ToTenantEntity(record)
	}
	return entities, nil
}

func (r *TenantRepository) Create(ctx context.Context, e *entity.Tenant) error {
	create := builder.BuildCreateTenant(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, tenantRepoName)
	}

	if created := mapper.ToTenantEntity(record); created != nil {
		*e = *created
	}
	return nil
}

func (r *TenantRepository) Update(ctx context.Context, e *entity.Tenant) error {
	update := builder.BuildUpdateTenant(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, tenantRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *TenantRepository) Delete(ctx context.Context, id int) error {
	if err := r.client.DB(ctx).Tenant.DeleteOneID(id).Exec(ctx); err != nil {
		return commonEnt.MapEntError(err, tenantRepoName)
	}
	return nil
}

func (r *TenantRepository) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := r.client.DB(ctx).Tenant.Query().Where(tenant.ID(id)).Exist(ctx)
	return exists, commonEnt.MapEntError(err, tenantRepoName)
}

package db

import (
	"context"

	commonEnt "go-link/common/pkg/database/ent"
	d "go-link/common/pkg/dto"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"

	"entgo.io/ent/dialect/sql"

	"go-link/identity/internal/adapters/driven/db/ent/builder"
	"go-link/identity/internal/adapters/driven/db/ent/generate/tenantmember"
	"go-link/identity/internal/adapters/driven/db/mapper"
	"go-link/identity/internal/core/entity"
	"go-link/identity/internal/ports"
)

const tenantMemberRepoName = "TenantMemberRepository"

type TenantMemberRepository struct {
	client *dbEnt.EntClient
}

func NewTenantMemberRepository(client *dbEnt.EntClient) ports.TenantMemberRepository {
	return &TenantMemberRepository{client: client}
}

func (r *TenantMemberRepository) Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.TenantMember], error) {
	client := r.client.DB(ctx)

	query := client.TenantMember.Query()
	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplyFilters(opts.Filters, s)
		})
	}

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantMemberRepoName)
	}

	if opts != nil {
		query.Where(func(s *sql.Selector) {
			commonEnt.ApplySort(opts.Sort, s)
			commonEnt.ApplyPagination(opts.Pagination, s)
		})
	}

	records, err := query.All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantMemberRepoName)
	}

	entities := make([]*entity.TenantMember, len(records))
	for i, record := range records {
		entities[i] = mapper.ToTenantMemberEntity(record)
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

	return &d.Paginated[*entity.TenantMember]{
		Records:    &entities,
		Pagination: meta,
	}, nil
}

func (r *TenantMemberRepository) Get(ctx context.Context, id int) (*entity.TenantMember, error) {
	record, err := r.client.DB(ctx).TenantMember.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantMemberRepoName)
	}
	return mapper.ToTenantMemberEntity(record), nil
}

func (r *TenantMemberRepository) Create(ctx context.Context, e *entity.TenantMember) error {
	create := builder.BuildCreateTenantMember(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, tenantMemberRepoName)
	}

	if created := mapper.ToTenantMemberEntity(record); created != nil {
		*e = *created
	}
	return nil
}

func (r *TenantMemberRepository) Update(ctx context.Context, e *entity.TenantMember) error {
	update := builder.BuildUpdateTenantMember(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, tenantMemberRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *TenantMemberRepository) Delete(ctx context.Context, id int) error {
	return commonEnt.MapEntError(r.client.DB(ctx).TenantMember.DeleteOneID(id).Exec(ctx), tenantMemberRepoName)
}

func (r *TenantMemberRepository) Exists(ctx context.Context, id int) (bool, error) {
	exists, err := r.client.DB(ctx).TenantMember.Query().Where(tenantmember.ID(id)).Exist(ctx)
	return exists, commonEnt.MapEntError(err, tenantMemberRepoName)
}

func (r *TenantMemberRepository) GetByUserAndTenant(ctx context.Context, userID, tenantID int) (*entity.TenantMember, error) {
	record, err := r.client.DB(ctx).TenantMember.Query().
		Where(tenantmember.UserID(userID), tenantmember.TenantID(tenantID)).
		Only(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, tenantMemberRepoName)
	}
	return mapper.ToTenantMemberEntity(record), nil
}

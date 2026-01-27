package db

import (
	"context"

	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	"go-link/billing/internal/adapters/driven/db/ent/builder"
	"go-link/billing/internal/adapters/driven/db/ent/generate/plan"
	"go-link/billing/internal/adapters/driven/db/mapper"
	"go-link/billing/internal/core/entity"
	"go-link/billing/internal/ports"
	commonEnt "go-link/common/pkg/database/ent"
)

const planRepoName = "PlanRepository"

type PlanRepository struct {
	client *dbEnt.EntClient
}

func NewPlanRepository(client *dbEnt.EntClient) ports.PlanRepository {
	return &PlanRepository{client: client}
}

func (r *PlanRepository) Get(ctx context.Context, id int) (*entity.Plan, error) {
	record, err := r.client.DB(ctx).Plan.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, planRepoName)
	}
	return mapper.ToPlanEntity(record), nil
}

func (r *PlanRepository) Create(ctx context.Context, e *entity.Plan) error {
	create := builder.BuildCreatePlan(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, planRepoName)
	}

	if created := mapper.ToPlanEntity(record); created != nil {
		*e = *created
	}
	return nil
}

func (r *PlanRepository) Update(ctx context.Context, e *entity.Plan) error {
	update := builder.BuildUpdatePlan(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, planRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *PlanRepository) Delete(ctx context.Context, id int) error {
	return commonEnt.MapEntError(r.client.DB(ctx).Plan.DeleteOneID(id).Exec(ctx), planRepoName)
}

func (r *PlanRepository) FindAll(ctx context.Context) ([]*entity.Plan, error) {
	records, err := r.client.DB(ctx).Plan.Query().All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, planRepoName)
	}

	entities := make([]*entity.Plan, len(records))
	for i, record := range records {
		entities[i] = mapper.ToPlanEntity(record)
	}
	return entities, nil
}

func (r *PlanRepository) FindActive(ctx context.Context) ([]*entity.Plan, error) {
	records, err := r.client.DB(ctx).Plan.Query().
		Where(plan.IsActive(true)).
		All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, planRepoName)
	}

	entities := make([]*entity.Plan, len(records))
	for i, record := range records {
		entities[i] = mapper.ToPlanEntity(record)
	}
	return entities, nil
}

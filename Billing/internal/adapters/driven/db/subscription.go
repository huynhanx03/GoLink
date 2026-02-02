package db

import (
	"context"

	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	"go-link/billing/internal/adapters/driven/db/ent/builder"
	entSub "go-link/billing/internal/adapters/driven/db/ent/generate/subscription"
	"go-link/billing/internal/adapters/driven/db/mapper"
	"go-link/billing/internal/core/entity"
	"go-link/billing/internal/ports"
	commonEnt "go-link/common/pkg/database/ent"
)

const subscriptionRepoName = "SubscriptionRepository"

type SubscriptionRepository struct {
	client *dbEnt.EntClient
}

func NewSubscriptionRepository(client *dbEnt.EntClient) ports.SubscriptionRepository {
	return &SubscriptionRepository{client: client}
}

func (r *SubscriptionRepository) Get(ctx context.Context, id int) (*entity.Subscription, error) {
	record, err := r.client.DB(ctx).Subscription.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, subscriptionRepoName)
	}
	return mapper.ToSubscriptionEntity(record), nil
}

func (r *SubscriptionRepository) GetByTenantID(ctx context.Context, tenantID int) (*entity.Subscription, error) {
	record, err := r.client.DB(ctx).Subscription.Query().
		Where(entSub.TenantID(tenantID)).
		Only(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, subscriptionRepoName)
	}
	return mapper.ToSubscriptionEntity(record), nil
}

func (r *SubscriptionRepository) Create(ctx context.Context, e *entity.Subscription) error {
	create := builder.BuildCreateSubscription(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, subscriptionRepoName)
	}

	if created := mapper.ToSubscriptionEntity(record); created != nil {
		*e = *created
	}
	return nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, e *entity.Subscription) error {
	update := builder.BuildUpdateSubscription(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, subscriptionRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id int) error {
	return commonEnt.MapEntError(r.client.DB(ctx).Subscription.DeleteOneID(id).Exec(ctx), subscriptionRepoName)
}

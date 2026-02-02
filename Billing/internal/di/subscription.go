package di

import (
	db "go-link/billing/internal/adapters/driven/db"
	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	driverHttp "go-link/billing/internal/adapters/driver/http"
	"go-link/billing/internal/core/service"
	"go-link/billing/internal/ports"
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/logger"
)

// SubscriptionContainer holds subscription-related dependencies.
type SubscriptionContainer struct {
	Repository ports.SubscriptionRepository
	Service    ports.SubscriptionService
	Handler    driverHttp.SubscriptionHandler
}

// InitSubscriptionDependencies initializes subscription dependencies.
func InitSubscriptionDependencies(
	client *dbEnt.EntClient,
	planRepo ports.PlanRepository,
	cache cache.CacheEngine,
	log *logger.LoggerZap,
) SubscriptionContainer {
	repository := db.NewSubscriptionRepository(client)
	service := service.NewSubscriptionService(repository, planRepo, cache, log)
	handler := driverHttp.NewSubscriptionHandler(service)

	return SubscriptionContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

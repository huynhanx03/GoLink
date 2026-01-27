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

// PlanContainer holds plan-related dependencies.
type PlanContainer struct {
	Repository ports.PlanRepository
	Service    ports.PlanService
	Handler    driverHttp.PlanHandler
}

// InitPlanDependencies initializes plan dependencies.
func InitPlanDependencies(
	client *dbEnt.EntClient,
	cache cache.LocalCache[string, any],
	log *logger.LoggerZap,
) PlanContainer {
	repository := db.NewPlanRepository(client)
	service := service.NewPlanService(repository, cache, log)
	handler := driverHttp.NewPlanHandler(service)

	return PlanContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

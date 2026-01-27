package di

import (
	"go-link/common/pkg/common/cache"
	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// TenantContainer holds tenant-related dependencies.
type TenantContainer struct {
	Repository ports.TenantRepository
	Service    ports.TenantService
	Handler    driverHttp.TenantHandler
}

// InitTenantDependencies initializes tenant dependencies.
func InitTenantDependencies(
	client *dbEnt.EntClient,
	cache cache.LocalCache[string, any],
) TenantContainer {
	repository := db.NewTenantRepository(client)
	service := service.NewTenantService(repository, cache)
	handler := driverHttp.NewTenantHandler(service)

	return TenantContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

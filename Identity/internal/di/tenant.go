package di

import (
	"go-link/common/pkg/common/cache"
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
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
	client *generate.Client,
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

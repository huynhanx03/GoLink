package di

import (
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
func InitTenantDependencies(client *generate.Client) TenantContainer {
	repository := db.NewTenantRepository(client)
	service := service.NewTenantService(repository)
	handler := driverHttp.NewTenantHandler(service)

	return TenantContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

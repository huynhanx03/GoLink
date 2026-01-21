package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// ResourceContainer holds resource-related dependencies.
type ResourceContainer struct {
	Repository ports.ResourceRepository
	Service    ports.ResourceService
	Handler    driverHttp.ResourceHandler
}

// InitResourceDependencies initializes resource dependencies.
func InitResourceDependencies(
	client *generate.Client,
	cacheService ports.CacheService,
) ResourceContainer {
	repository := db.NewResourceRepository(client)
	service := service.NewResourceService(repository, cacheService)
	handler := driverHttp.NewResourceHandler(service)

	return ResourceContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

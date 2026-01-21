package di

import (
	"go-link/common/pkg/common/cache"
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// DomainContainer holds domain-related dependencies.
type DomainContainer struct {
	Repository ports.DomainRepository
	Service    ports.DomainService
	Handler    driverHttp.DomainHandler
}

// InitDomainDependencies initializes domain dependencies.
func InitDomainDependencies(
	client *generate.Client,
	cache cache.LocalCache[string, any],
) DomainContainer {
	repository := db.NewDomainRepository(client)
	service := service.NewDomainService(repository, cache)
	handler := driverHttp.NewDomainHandler(service)

	return DomainContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

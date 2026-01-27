package di

import (
	"go-link/common/pkg/common/cache"

	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// AttributeDefinitionContainer holds attribute definition related dependencies.
type AttributeDefinitionContainer struct {
	Repository ports.AttributeDefinitionRepository
	Service    ports.AttributeDefinitionService
	Handler    driverHttp.AttributeDefinitionHandler
}

// InitAttributeDefinitionDependencies initializes attribute definition dependencies.
func InitAttributeDefinitionDependencies(
	client *dbEnt.EntClient,
	cache cache.LocalCache[string, any],
) AttributeDefinitionContainer {
	repository := db.NewAttributeDefinitionRepository(client)
	service := service.NewAttributeDefinitionService(repository, cache)
	handler := driverHttp.NewAttributeDefinitionHandler(service)

	return AttributeDefinitionContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

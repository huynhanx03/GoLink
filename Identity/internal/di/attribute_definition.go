package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
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
func InitAttributeDefinitionDependencies(client *generate.Client) AttributeDefinitionContainer {
	repository := db.NewAttributeDefinitionRepository(client)
	service := service.NewAttributeDefinitionService(repository)
	handler := driverHttp.NewAttributeDefinitionHandler(service)

	return AttributeDefinitionContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// PermissionContainer holds permission-related dependencies.
type PermissionContainer struct {
	Repository ports.PermissionRepository
	Service    ports.PermissionService
	Handler    driverHttp.PermissionHandler
}

// InitPermissionDependencies initializes permission dependencies.
func InitPermissionDependencies(client *generate.Client) PermissionContainer {
	repository := db.NewPermissionRepository(client)
	service := service.NewPermissionService(repository)
	handler := driverHttp.NewPermissionHandler(service)

	return PermissionContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

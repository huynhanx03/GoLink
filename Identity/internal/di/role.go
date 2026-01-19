package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// RoleContainer holds role-related dependencies.
type RoleContainer struct {
	Repository ports.RoleRepository
	Service    ports.RoleService
	Handler    driverHttp.RoleHandler
}

// InitRoleDependencies initializes role dependencies.
func InitRoleDependencies(client *generate.Client) RoleContainer {
	repository := db.NewRoleRepository(client)
	service := service.NewRoleService(repository)
	handler := driverHttp.NewRoleHandler(service)

	return RoleContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

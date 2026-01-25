package di

import (
	db "go-link/identity/internal/adapters/driven/db"
	dbEnt "go-link/identity/internal/adapters/driven/db/ent"
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
func InitRoleDependencies(
	client *dbEnt.EntClient,
	cacheService ports.CacheService,
) RoleContainer {
	repository := db.NewRoleRepository(client)
	service := service.NewRoleService(repository, cacheService)
	handler := driverHttp.NewRoleHandler(service)

	return RoleContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

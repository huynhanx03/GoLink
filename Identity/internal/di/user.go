package di

import (
	"go-link/common/pkg/common/cache"
	userDB "go-link/identity/internal/adapters/driven/db"
	"go-link/identity/internal/adapters/driven/db/ent/generate"
	"go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

type UserContainer struct {
	Repository ports.UserRepository
	Service    ports.UserService
	Handler    http.UserHandler
}

func InitUserRepository(client *generate.Client) ports.UserRepository {
	return userDB.NewUserRepository(client)
}

func InitUserDependencies(
	client *generate.Client,
	repo ports.UserRepository,
	authService ports.AuthenticationService,
	credentialRepo ports.CredentialRepository,
	tenantRepo ports.TenantRepository,
	tenantMemberRepo ports.TenantMemberRepository,
	roleRepo ports.RoleRepository,
	attrDefRepo ports.AttributeDefinitionRepository,
	attrValueRepo ports.UserAttributeValueRepository,
	localCache cache.LocalCache[string, any],
) UserContainer {
	svc := service.NewUserService(
		repo,
		credentialRepo,
		tenantRepo,
		tenantMemberRepo,
		roleRepo,
		attrDefRepo,
		attrValueRepo,
		localCache,
	)
	handler := http.NewUserHandler(svc, authService)

	return UserContainer{
		Repository: repo,
		Service:    svc,
		Handler:    handler,
	}
}

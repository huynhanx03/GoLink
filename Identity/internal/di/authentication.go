package di

import (
	"go-link/common/pkg/common/cache"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// AuthenticationContainer holds authentication dependencies.
type AuthenticationContainer struct {
	Service ports.AuthenticationService
	Handler driverHttp.AuthenticationHandler
}

// InitAuthenticationDependencies initializes authentication dependencies.
func InitAuthenticationDependencies(
	userRepo ports.UserRepository,
	credentialRepo ports.CredentialRepository,
	tenantRepo ports.TenantRepository,
	tenantMemberRepo ports.TenantMemberRepository,
	roleRepo ports.RoleRepository,
	permissionRepo ports.PermissionRepository,
	resourceRepo ports.ResourceRepository,
	attrDefinitionRepo ports.AttributeDefinitionRepository,
	attrValueRepo ports.UserAttributeValueRepository,
	cache cache.LocalCache[string, any],
	cacheService ports.CacheService,
) AuthenticationContainer {
	service := service.NewAuthenticationService(
		userRepo,
		credentialRepo,
		tenantRepo,
		tenantMemberRepo,
		roleRepo,
		permissionRepo,
		resourceRepo,
		attrDefinitionRepo,
		attrValueRepo,
		cache,
		cacheService,
	)
	handler := driverHttp.NewAuthenticationHandler(service)

	return AuthenticationContainer{
		Service: service,
		Handler: handler,
	}
}

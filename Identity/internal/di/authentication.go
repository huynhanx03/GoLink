package di

import (
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/mq/kafka"
	"go-link/identity/global"
	driverHttp "go-link/identity/internal/adapters/driver/http"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
	"go-link/identity/pkg/oauth"
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
	fedIdentityRepo ports.FederatedIdentityRepository,
	cache cache.LocalCache[string, any],
	cacheService ports.CacheService,
	producer kafka.SyncProducer,
) AuthenticationContainer {
	oauthProviders := map[string]oauth.Provider{
		"google": oauth.NewGoogleProvider(
			global.Config.Google.ClientID,
			global.Config.Google.ClientSecret,
			global.Config.Google.RedirectURL,
		),
	}

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
		fedIdentityRepo,
		oauthProviders,
		cache,
		cacheService,
		producer,
	)
	handler := driverHttp.NewAuthenticationHandler(service)

	return AuthenticationContainer{
		Service: service,
		Handler: handler,
	}
}

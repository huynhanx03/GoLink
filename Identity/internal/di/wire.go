package di

import "go-link/identity/global"

// SetupDependencies initializes all dependencies and returns the container.
func SetupDependencies() *Container {
	client := global.EntClient

	container := &Container{
		// Full containers (with Handler)
		TenantContainer:     InitTenantDependencies(client),
		RoleContainer:       InitRoleDependencies(client),
		PermissionContainer: InitPermissionDependencies(client),
		DomainContainer:     InitDomainDependencies(client),
		ResourceContainer:   InitResourceDependencies(client),

		// Repository-only containers
		UserContainer:                InitUserDependencies(client),
		CredentialContainer:          InitCredentialDependencies(client),
		FederatedIdentityContainer:   InitFederatedIdentityDependencies(client),
		TenantMemberContainer:        InitTenantMemberDependencies(client),
		AttributeDefinitionContainer: InitAttributeDefinitionDependencies(client),
		UserAttributeValueContainer:  InitUserAttributeValueDependencies(client),
	}

	GlobalContainer = container
	return container
}

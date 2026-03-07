package di

import (
	"fmt"

	"go-link/common/pkg/mq/kafka"
	"go-link/identity/global"
)

// SetupDependencies initializes all dependencies and returns the container.
func SetupDependencies() *Container {
	client := global.EntClient

	credentialContainer := InitCredentialDependencies(client)
	tenantMemberContainer := InitTenantMemberDependencies(client)
	cacheContainer := InitCacheDependencies(global.Tinylfu)
	roleContainer := InitRoleDependencies(client, cacheContainer.Service)
	attrValueContainer := InitUserAttributeValueDependencies(client)
	attrDefinitionContainer := InitAttributeDefinitionDependencies(client, global.Tinylfu)
	tenantContainer := InitTenantDependencies(client, tenantMemberContainer.Repository, global.Tinylfu)
	permissionContainer := InitPermissionDependencies(client, cacheContainer.Service)
	resourceContainer := InitResourceDependencies(client, cacheContainer.Service)
	domainContainer := InitDomainDependencies(client, global.Tinylfu)
	federatedIdentityContainer := InitFederatedIdentityDependencies(client)

	userRepo := InitUserRepository(client)

	// Create Kafka SyncProducer for dispatching notification events.
	// A failure here is non-fatal: the service continues without notification dispatch.
	kafkaCfg := &kafka.Config{
		Brokers:  global.Config.Kafka.Brokers,
		ClientID: "identity-service",
	}
	producer, err := kafka.NewSyncProducer(kafkaCfg)
	if err != nil {
		fmt.Printf("[WARN] Failed to create Kafka producer, notifications disabled: %v\n", err)
	}

	authContainer := InitAuthenticationDependencies(
		userRepo,
		credentialContainer.Repository,
		tenantContainer.Repository,
		tenantMemberContainer.Repository,
		roleContainer.Repository,
		permissionContainer.Repository,
		resourceContainer.Repository,
		attrDefinitionContainer.Repository,
		attrValueContainer.Repository,
		federatedIdentityContainer.Repository,

		global.Tinylfu,
		cacheContainer.Service,
		producer,
	)

	userContainer := InitUserDependencies(
		client,
		userRepo,
		authContainer.Service,
		credentialContainer.Repository,
		tenantContainer.Repository,
		tenantMemberContainer.Repository,
		roleContainer.Repository,
		attrDefinitionContainer.Repository,
		attrValueContainer.Repository,
		global.Tinylfu,
	)

	container := &Container{
		TenantContainer:              tenantContainer,
		RoleContainer:                roleContainer,
		PermissionContainer:          permissionContainer,
		DomainContainer:              domainContainer,
		ResourceContainer:            resourceContainer,
		AuthenticationContainer:      authContainer,
		UserContainer:                userContainer,
		CredentialContainer:          credentialContainer,
		FederatedIdentityContainer:   federatedIdentityContainer,
		TenantMemberContainer:        tenantMemberContainer,
		AttributeDefinitionContainer: attrDefinitionContainer,
		UserAttributeValueContainer:  attrValueContainer,
	}

	GlobalContainer = container
	return container
}

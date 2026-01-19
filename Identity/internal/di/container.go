package di

// Container holds all dependency containers for the Identity service.
type Container struct {
	TenantContainer              TenantContainer
	RoleContainer                RoleContainer
	PermissionContainer          PermissionContainer
	DomainContainer              DomainContainer
	ResourceContainer            ResourceContainer
	UserContainer                UserContainer
	CredentialContainer          CredentialContainer
	FederatedIdentityContainer   FederatedIdentityContainer
	TenantMemberContainer        TenantMemberContainer
	AttributeDefinitionContainer AttributeDefinitionContainer
	UserAttributeValueContainer  UserAttributeValueContainer
}

// GlobalContainer is the global instance of Container.
var GlobalContainer *Container

package permissions

// Resource Keys
const (
	ResourceKeyGeneration          = "generation"
	ResourceKeyDomain              = "domains"
	ResourceKeyTenant              = "tenant"
	ResourceKeyUser                = "user"
	ResourceKeyRole                = "role"
	ResourceKeyPermission          = "permission"
	ResourceKeyResource            = "resource"
	ResourceKeyAttributeDefinition = "attribute_definition"
	ResourceKeyBilling             = "billing"
	ResourceKeyPayment             = "payment"
)

// Permission Scopes (Bitmask)
const (
	PermissionScopeCreate = 1 // 0001
	PermissionScopeRead   = 2 // 0010
	PermissionScopeUpdate = 4 // 0100
	PermissionScopeDelete = 8 // 1000
)

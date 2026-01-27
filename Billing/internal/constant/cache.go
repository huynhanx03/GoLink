package constant

import "time"

const (
	// Cache Prefix
	CacheKeyPrefixPlanID         = "billing:plan:id:"
	CacheKeyPrefixInvoiceID      = "billing:invoice:id:"
	CacheKeyPrefixSubscriptionID = "billing:subscription:id:"

	// Cache Config
	CacheCostID     = 1
	CacheTTLDefault = time.Hour * 1
)

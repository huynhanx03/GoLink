package constant

import "time"

const (
	LinkCachePrefix   = "link::"
	LinkCacheTTL      = 1 * time.Hour
	UserLevelCacheTTL = 1 * time.Hour

	RedisKeyUsageTenantLinks = "usage:tenant:%d:links"
	RedisKeyUserLevel        = "sys:user:%d:level"
	LocalCacheKeyTierConfig  = "config:tier:%d:max_links"

	CacheCostQuota = 1
)

package di

import (
	"go-link/common/pkg/common/cache"
	"go-link/identity/internal/core/service"
	"go-link/identity/internal/ports"
)

// CacheContainer holds cache-related dependencies.
type CacheContainer struct {
	Service ports.CacheService
}

// InitCacheDependencies initializes cache dependencies.
func InitCacheDependencies(
	cache cache.LocalCache[string, any],
) CacheContainer {
	svc := service.NewCacheService(cache)

	return CacheContainer{
		Service: svc,
	}
}

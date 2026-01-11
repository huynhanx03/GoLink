package cache

import (
	"context"

	"go-link/common/pkg/common/cache"

	"go-link/generation/internal/constant"
	"go-link/generation/internal/core/entity"
	"go-link/generation/internal/ports"
)

type linkCache struct {
	redis cache.CacheEngine
}

func NewLink(redis cache.CacheEngine) ports.LinkCacheRepository {
	return &linkCache{
		redis: redis,
	}
}

func (l *linkCache) getKey(id string) string {
	return constant.LinkCachePrefix + id
}

func (l *linkCache) Set(ctx context.Context, link *entity.Link) error {
	return cache.HandleSetCache(ctx, link, l.redis, l.getKey(link.ID), constant.LinkCacheTTL)
}
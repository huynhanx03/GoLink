package cache

import (
	"context"

	"go-link/common/pkg/common/cache"

	"go-link/redirection/internal/constant"
	"go-link/redirection/internal/core/entity"
	"go-link/redirection/internal/ports"
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

func (l *linkCache) Get(ctx context.Context, id string) (*entity.Link, error) {
	var link entity.Link

	if err := cache.HandleHitCache(ctx, &link, l.redis, l.getKey(id)); err != nil {
		return nil, err
	}

	return &link, nil
}

func (l *linkCache) DeleteBatch(ctx context.Context, ids []string) error {
	idKeys := make([]string, len(ids))
	for i, id := range ids {
		idKeys[i] = l.getKey(id)
	}
	return l.redis.DeleteBatch(ctx, idKeys)
}

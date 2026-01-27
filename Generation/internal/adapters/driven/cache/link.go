package cache

import (
	"context"
	"fmt"

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

func (l *linkCache) IncrementQuota(ctx context.Context, tenantID int) (int64, error) {
	key := fmt.Sprintf(constant.RedisKeyUsageTenantLinks, tenantID)
	return l.redis.Incr(ctx, key)
}

func (l *linkCache) DecrementQuota(ctx context.Context, tenantID int) (int64, error) {
	key := fmt.Sprintf(constant.RedisKeyUsageTenantLinks, tenantID)
	return l.redis.Decr(ctx, key)
}

func (l *linkCache) GetUserLevel(ctx context.Context, userID int) (int, error) {
	key := fmt.Sprintf(constant.RedisKeyUserLevel, userID)
	var level int
	if err := cache.HandleHitCache(ctx, &level, l.redis, key); err != nil {
		return 0, err
	}
	return level, nil
}

func (l *linkCache) SetUserLevel(ctx context.Context, userID int, level int) error {
	key := fmt.Sprintf(constant.RedisKeyUserLevel, userID)
	return cache.HandleSetCache(ctx, level, l.redis, key, constant.UserLevelCacheTTL)
}

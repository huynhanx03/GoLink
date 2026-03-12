package cache

import (
	"context"
	"time"

	"go-link/common/pkg/common/cache"
	"go-link/notification/internal/constant"
	"go-link/notification/internal/ports"
)

const idempotencyTTL = 24 * time.Hour

type redisIdempotencyChecker struct {
	engine cache.CacheEngine
}

// NewRedisIdempotencyChecker creates a Redis-backed IdempotencyChecker.
func NewRedisIdempotencyChecker(engine cache.CacheEngine) ports.IdempotencyChecker {
	return &redisIdempotencyChecker{engine: engine}
}

// TryAcquire atomically sets the idempotency key via SETNX.
// Returns true if the key was newly set (first caller), false if already exists (duplicate).
// On Redis error, fails open (returns true) to avoid blocking legitimate notifications.
func (c *redisIdempotencyChecker) TryAcquire(ctx context.Context, key string) (bool, error) {
	cacheKey := constant.CacheKeyPrefixIdempotency + key
	acquired, err := c.engine.SetNX(ctx, cacheKey, "1", idempotencyTTL)
	if err != nil {
		return true, nil // Fail open: allow processing on cache errors.
	}
	return acquired, nil
}

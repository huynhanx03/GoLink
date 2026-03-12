package service

import (
	"context"
	"fmt"
	"time"

	"go-link/common/pkg/common/cache"
)

const rateLimitPrefix = "notification:ratelimit:"

// RateLimiter implements a fixed-window counter rate limiter backed by Redis.
type RateLimiter struct {
	engine cache.CacheEngine
}

// NewRateLimiter creates a new RateLimiter using the provided cache engine.
func NewRateLimiter(engine cache.CacheEngine) *RateLimiter {
	return &RateLimiter{engine: engine}
}

// Allow checks whether the user has not exceeded maxPerWindow notifications for the given channel
// within the sliding window duration. Returns true if allowed, false if rate-limited.
// On Redis error, fails open (returns true) to avoid blocking legitimate notifications.
func (r *RateLimiter) Allow(ctx context.Context, userID, channel string, maxPerWindow int64, window time.Duration) (bool, error) {
	key := fmt.Sprintf("%s%s:%s", rateLimitPrefix, userID, channel)
	now := time.Now()
	nowMs := float64(now.UnixMilli())
	windowMs := float64(window.Milliseconds())
	minScore := fmt.Sprintf("%f", nowMs-windowMs)

	// 1. Remove entries older than the window
	_ = r.engine.ZRemRangeByScore(ctx, key, "-inf", minScore)

	// 2. Count current entries
	count, err := r.engine.ZCount(ctx, key, "-inf", "+inf")
	if err != nil {
		return true, nil // fail open
	}

	if count >= maxPerWindow {
		return false, nil
	}

	// 3. Add new attempt
	err = r.engine.ZAdd(ctx, key, &cache.ZMember{
		Score:  nowMs,
		Member: nowMs, // Unique enough for millisecond precision
	})
	if err != nil {
		return true, nil // fail open
	}

	// 4. Set TTL to ensure key cleanup eventually (max window + some buffer)
	_ = r.engine.Expire(ctx, key, window*2)

	return true, nil
}

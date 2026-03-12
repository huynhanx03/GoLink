package cache

import (
	"context"
	"encoding/binary"
	"strconv"

	"go-link/common/pkg/common/cache"
	"go-link/notification/internal/constant"
)

// UnreadCounter is a Redis-backed counter for per-user unread notification counts.
type UnreadCounter struct {
	engine cache.CacheEngine
}

// NewUnreadCounter creates a new Redis-backed UnreadCounter.
func NewUnreadCounter(engine cache.CacheEngine) *UnreadCounter {
	return &UnreadCounter{engine: engine}
}

// cacheKey returns the Redis key for a given user's unread count.
func (c *UnreadCounter) cacheKey(userID string) string {
	return constant.CacheKeyPrefixUnreadCount + userID
}

// Increment increments the unread counter by 1 for the given user.
func (c *UnreadCounter) Increment(ctx context.Context, userID string) error {
	_, err := c.engine.Incr(ctx, c.cacheKey(userID))
	return err
}

// Decrement decrements the unread counter by 1 for the given user (floor at 0).
func (c *UnreadCounter) Decrement(ctx context.Context, userID string) error {
	_, err := c.engine.Decr(ctx, c.cacheKey(userID))
	return err
}

// Get returns the current unread count for the given user.
// Returns 0 if the key does not exist or on any error (fail-open).
func (c *UnreadCounter) Get(ctx context.Context, userID string) (int64, error) {
	raw, found, err := c.engine.Get(ctx, c.cacheKey(userID))
	if err != nil || !found {
		return 0, nil
	}

	// Redis returns the integer as a decimal string (e.g. "42").
	// Try string parse first, fall back to little-endian binary.
	if n, err := strconv.ParseInt(string(raw), 10, 64); err == nil {
		return n, nil
	}

	if len(raw) == 8 {
		return int64(binary.LittleEndian.Uint64(raw)), nil
	}

	return 0, nil
}

// Reset removes the unread counter key for the given user.
func (c *UnreadCounter) Reset(ctx context.Context, userID string) error {
	return c.engine.Delete(ctx, c.cacheKey(userID))
}

package service

import (
	"context"
	"fmt"
	"time"

	"go-link/common/pkg/common/cache"
	"go-link/notification/internal/core/entity"
	"go-link/notification/internal/ports"
)

const digestKeyPrefix = "notification:digest:"

// DigestService handles notification aggregation using Redis.
type DigestService struct {
	cache cache.CacheEngine
}

func NewDigestService(cache cache.CacheEngine) ports.DigestService {
	return &DigestService{cache: cache}
}

// ShouldDigest checks if a notification should be aggregated.
func (s *DigestService) ShouldDigest(notification *entity.Notification) bool {
	return notification.CollapseKey != ""
}

// AddToDigest stores notification ID in a sorted set for later batching.
func (s *DigestService) AddToDigest(ctx context.Context, n *entity.Notification) error {
	if n.CollapseKey == "" {
		return nil
	}

	key := fmt.Sprintf("%s%s:%s", digestKeyPrefix, n.Recipient.UserID, n.CollapseKey)
	member := &cache.ZMember{
		Score:  float64(time.Now().UnixMilli()),
		Member: n.ID,
	}

	if err := s.cache.ZAdd(ctx, key, member); err != nil {
		return err
	}

	return s.cache.Expire(ctx, key, 24*time.Hour)
}

// ScanPendingDigests scans for all keys matching the digest prefix.
func (s *DigestService) ScanPendingDigests(ctx context.Context) ([]string, error) {
	return s.cache.Keys(ctx, digestKeyPrefix+"*")
}

// ConsumeDigest returns all IDs in the digest and deletes the key.
func (s *DigestService) ConsumeDigest(ctx context.Context, key string) ([]string, error) {
	ids, err := s.cache.ZRange(ctx, key, 0, -1)
	if err != nil || len(ids) == 0 {
		return ids, err
	}

	if err := s.cache.Delete(ctx, key); err != nil {
		return nil, err
	}

	return ids, nil
}

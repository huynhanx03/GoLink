package ports

import "context"

// UnreadCounter defines the contract for managing per-user unread notification counts.
type UnreadCounter interface {
	// Increment increments the unread counter by 1 for the given user.
	Increment(ctx context.Context, userID string) error

	// Decrement decrements the unread counter by 1 for the given user (floor at 0).
	Decrement(ctx context.Context, userID string) error

	// Get returns the current unread count for the given user.
	Get(ctx context.Context, userID string) (int64, error)

	// Reset removes the unread counter key for the given user.
	Reset(ctx context.Context, userID string) error
}

package ports

import "context"

// IdempotencyChecker defines the contract for idempotency checking.
type IdempotencyChecker interface {
	// TryAcquire atomically checks and marks a key as processed.
	// Returns true if this is the first call (key was newly set), false if already processed.
	TryAcquire(ctx context.Context, key string) (bool, error)
}

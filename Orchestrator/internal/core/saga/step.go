package saga

import "context"

// Step defines a single step in a Saga.
// Each step must implement Execute (forward action) and Compensate (rollback action).
type Step interface {
	Name() string
	Execute(ctx context.Context) error
	Compensate(ctx context.Context) error
}

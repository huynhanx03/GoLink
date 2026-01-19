package ports

import (
	"context"

	"go-link/identity/internal/core/entity"
)

// UserRepository defines the interface for user persistence.
type UserRepository interface {
	Get(ctx context.Context, id int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	Create(ctx context.Context, e *entity.User) error
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
}

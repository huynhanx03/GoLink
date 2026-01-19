package ports

import (
	"context"

	d "go-link/common/pkg/dto"

	"go-link/identity/internal/core/entity"
)

// TenantMemberRepository defines the tenant member data access interface.
type TenantMemberRepository interface {
	Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.TenantMember], error)
	Get(ctx context.Context, id int) (*entity.TenantMember, error)
	Create(ctx context.Context, e *entity.TenantMember) error
	Update(ctx context.Context, e *entity.TenantMember) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) (bool, error)
}

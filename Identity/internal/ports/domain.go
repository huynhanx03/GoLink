package ports

import (
	"context"

	d "go-link/common/pkg/dto"

	"go-link/identity/internal/core/dto"
	"go-link/identity/internal/core/entity"
)

// DomainRepository defines the domain data access interface.
type DomainRepository interface {
	Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*entity.Domain], error)
	Get(ctx context.Context, id int) (*entity.Domain, error)
	Create(ctx context.Context, e *entity.Domain) error
	Update(ctx context.Context, e *entity.Domain) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) (bool, error)
}

// DomainService defines the domain business logic interface.
type DomainService interface {
	Find(ctx context.Context, opts *d.QueryOptions) (*d.Paginated[*dto.DomainResponse], error)
	Get(ctx context.Context, id int) (*dto.DomainResponse, error)
	Create(ctx context.Context, req *dto.CreateDomainRequest) (*dto.DomainResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdateDomainRequest) (*dto.DomainResponse, error)
	Delete(ctx context.Context, id int) error
}

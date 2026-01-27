package ports

import (
	"context"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
)

// InvoiceRepository defines the invoice data access interface.
type InvoiceRepository interface {
	Get(ctx context.Context, id int) (*entity.Invoice, error)
	Create(ctx context.Context, e *entity.Invoice) error
	Update(ctx context.Context, e *entity.Invoice) error
	Delete(ctx context.Context, id int) error
	FindByTenantID(ctx context.Context, tenantID int) ([]*entity.Invoice, error)
}

// InvoiceService defines the invoice business logic interface.
type InvoiceService interface {
	Get(ctx context.Context, id int) (*dto.InvoiceResponse, error)
	Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error)
	Update(ctx context.Context, id int, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error)
	Delete(ctx context.Context, id int) error
	FindMine(ctx context.Context, tenantID int) ([]*dto.InvoiceResponse, error)
}

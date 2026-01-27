package http

import (
	"context"

	"go-link/common/pkg/common/http/handler"
	"go-link/common/pkg/constraints"

	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/ports"
)

// InvoiceHandler defines the invoice HTTP handler interface.
type InvoiceHandler interface {
	Get(ctx context.Context, req *dto.GetInvoiceRequest) (*dto.InvoiceResponse, error)
	Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error)
	Update(ctx context.Context, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error)
	Delete(ctx context.Context, req *dto.DeleteInvoiceRequest) (*dto.InvoiceResponse, error)
	FindMine(ctx context.Context, req *dto.FindMyInvoicesRequest) ([]*dto.InvoiceResponse, error)
}

type invoiceHandler struct {
	handler.BaseHandler
	invoiceService ports.InvoiceService
}

// NewInvoiceHandler creates a new InvoiceHandler instance.
func NewInvoiceHandler(invoiceService ports.InvoiceService) InvoiceHandler {
	return &invoiceHandler{
		invoiceService: invoiceService,
	}
}

// Get retrieves an invoice by ID.
func (h *invoiceHandler) Get(ctx context.Context, req *dto.GetInvoiceRequest) (*dto.InvoiceResponse, error) {
	return h.invoiceService.Get(ctx, req.ID)
}

// Create creates a new invoice.
func (h *invoiceHandler) Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	return h.invoiceService.Create(ctx, req)
}

// Update updates an existing invoice.
func (h *invoiceHandler) Update(ctx context.Context, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error) {
	return h.invoiceService.Update(ctx, req.ID, req)
}

// Delete removes an invoice by ID.
func (h *invoiceHandler) Delete(ctx context.Context, req *dto.DeleteInvoiceRequest) (*dto.InvoiceResponse, error) {
	return nil, h.invoiceService.Delete(ctx, req.ID)
}

// FindMine retrieves all invoices for the current tenant.
func (h *invoiceHandler) FindMine(ctx context.Context, req *dto.FindMyInvoicesRequest) ([]*dto.InvoiceResponse, error) {
	tenantID, ok := ctx.Value(constraints.ContextKeyTenantID).(int)
	if !ok {
		return nil, nil
	}
	return h.invoiceService.FindMine(ctx, tenantID)
}

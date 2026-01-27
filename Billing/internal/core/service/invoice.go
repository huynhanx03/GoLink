package service

import (
	"context"
	"strconv"

	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/logger"

	"go-link/billing/internal/constant"
	"go-link/billing/internal/core/dto"
	"go-link/billing/internal/core/entity"
	"go-link/billing/internal/core/mapper"
	"go-link/billing/internal/ports"
)

const invoiceServiceName = "InvoiceService"

type invoiceService struct {
	invoiceRepo ports.InvoiceRepository
	cache       cache.CacheEngine
	log         *logger.LoggerZap
}

// NewInvoiceService creates a new InvoiceService instance.
func NewInvoiceService(
	invoiceRepo ports.InvoiceRepository,
	cache cache.CacheEngine,
	log *logger.LoggerZap,
) ports.InvoiceService {
	return &invoiceService{
		invoiceRepo: invoiceRepo,
		cache:       cache,
		log:         log,
	}
}

// Get retrieves an invoice by ID.
func (s *invoiceService) Get(ctx context.Context, id int) (*dto.InvoiceResponse, error) {
	cacheKey := constant.CacheKeyPrefixInvoiceID + strconv.Itoa(id)
	var invoice *entity.Invoice
	if err := cache.HandleHitCache(ctx, &invoice, s.cache, cacheKey); err == nil {
		return mapper.ToInvoiceResponse(invoice), nil
	}

	invoice, err := s.invoiceRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = cache.HandleSetCache(ctx, invoice, s.cache, cacheKey, constant.CacheTTLDefault)
	return mapper.ToInvoiceResponse(invoice), nil
}

// Create creates a new invoice.
func (s *invoiceService) Create(ctx context.Context, req *dto.CreateInvoiceRequest) (*dto.InvoiceResponse, error) {
	invoice := mapper.ToInvoiceEntityFromCreate(req)

	if err := s.invoiceRepo.Create(ctx, invoice); err != nil {
		return nil, err
	}

	return mapper.ToInvoiceResponse(invoice), nil
}

// Update updates an existing invoice.
func (s *invoiceService) Update(ctx context.Context, id int, req *dto.UpdateInvoiceRequest) (*dto.InvoiceResponse, error) {
	invoice, err := s.invoiceRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Status != nil {
		invoice.Status = *req.Status
	}
	if req.PaymentID != nil {
		invoice.PaymentID = req.PaymentID
	}

	if err := s.invoiceRepo.Update(ctx, invoice); err != nil {
		return nil, err
	}

	cache.HandleUpdateCache(ctx, invoice, s.cache, constant.CacheKeyPrefixInvoiceID+strconv.Itoa(id), constant.CacheTTLDefault)

	return mapper.ToInvoiceResponse(invoice), nil
}

// Delete removes an invoice by ID.
func (s *invoiceService) Delete(ctx context.Context, id int) error {
	if err := s.invoiceRepo.Delete(ctx, id); err != nil {
		return err
	}

	_ = cache.HandleDeleteCache(ctx, s.cache, constant.CacheKeyPrefixInvoiceID+strconv.Itoa(id))

	return nil
}

// FindMine retrieves all invoices for a tenant.
func (s *invoiceService) FindMine(ctx context.Context, tenantID int) ([]*dto.InvoiceResponse, error) {
	invoices, err := s.invoiceRepo.FindByTenantID(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.InvoiceResponse, 0, len(invoices))
	for _, i := range invoices {
		responses = append(responses, mapper.ToInvoiceResponse(i))
	}

	return responses, nil
}

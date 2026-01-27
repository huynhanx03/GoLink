package db

import (
	"context"

	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	"go-link/billing/internal/adapters/driven/db/ent/builder"
	"go-link/billing/internal/adapters/driven/db/ent/generate/invoice"
	"go-link/billing/internal/adapters/driven/db/mapper"
	"go-link/billing/internal/core/entity"
	"go-link/billing/internal/ports"
	commonEnt "go-link/common/pkg/database/ent"
)

const invoiceRepoName = "InvoiceRepository"

type InvoiceRepository struct {
	client *dbEnt.EntClient
}

func NewInvoiceRepository(client *dbEnt.EntClient) ports.InvoiceRepository {
	return &InvoiceRepository{client: client}
}

func (r *InvoiceRepository) Get(ctx context.Context, id int) (*entity.Invoice, error) {
	record, err := r.client.DB(ctx).Invoice.Get(ctx, id)
	if err != nil {
		return nil, commonEnt.MapEntError(err, invoiceRepoName)
	}
	return mapper.ToInvoiceEntity(record), nil
}

func (r *InvoiceRepository) Create(ctx context.Context, e *entity.Invoice) error {
	create := builder.BuildCreateInvoice(ctx, e)
	record, err := create.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, invoiceRepoName)
	}

	if created := mapper.ToInvoiceEntity(record); created != nil {
		*e = *created
	}
	return nil
}

func (r *InvoiceRepository) Update(ctx context.Context, e *entity.Invoice) error {
	update := builder.BuildUpdateInvoice(ctx, e)
	record, err := update.Save(ctx)
	if err != nil {
		return commonEnt.MapEntError(err, invoiceRepoName)
	}
	e.UpdatedAt = record.UpdatedAt
	return nil
}

func (r *InvoiceRepository) Delete(ctx context.Context, id int) error {
	return commonEnt.MapEntError(r.client.DB(ctx).Invoice.DeleteOneID(id).Exec(ctx), invoiceRepoName)
}

func (r *InvoiceRepository) FindByTenantID(ctx context.Context, tenantID int) ([]*entity.Invoice, error) {
	records, err := r.client.DB(ctx).Invoice.Query().
		Where(invoice.TenantID(tenantID)).
		All(ctx)
	if err != nil {
		return nil, commonEnt.MapEntError(err, invoiceRepoName)
	}

	entities := make([]*entity.Invoice, len(records))
	for i, record := range records {
		entities[i] = mapper.ToInvoiceEntity(record)
	}
	return entities, nil
}

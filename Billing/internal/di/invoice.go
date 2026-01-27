package di

import (
	db "go-link/billing/internal/adapters/driven/db"
	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
	driverHttp "go-link/billing/internal/adapters/driver/http"
	"go-link/billing/internal/core/service"
	"go-link/billing/internal/ports"
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/logger"
)

// InvoiceContainer holds invoice-related dependencies.
type InvoiceContainer struct {
	Repository ports.InvoiceRepository
	Service    ports.InvoiceService
	Handler    driverHttp.InvoiceHandler
}

// InitInvoiceDependencies initializes invoice dependencies.
func InitInvoiceDependencies(
	client *dbEnt.EntClient,
	cache cache.CacheEngine,
	log *logger.LoggerZap,
) InvoiceContainer {
	repository := db.NewInvoiceRepository(client)
	service := service.NewInvoiceService(repository, cache, log)
	handler := driverHttp.NewInvoiceHandler(service)

	return InvoiceContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

package di

import (
	db "go-link/generation/internal/adapters/driven/db"
	driverHttp "go-link/generation/internal/adapters/driver/http"
	"go-link/generation/internal/core/service"
	"go-link/generation/internal/ports"
)

type LinkContainer struct {
	Repository ports.LinkRepository
	Service    ports.LinkService
	Handler    driverHttp.LinkHandler
}

func InitLinkDependencies() LinkContainer {
	// Repository
	repository := db.NewLinkRepository()

	// Service
	service := service.NewLinkService(repository)

	// Handler
	handler := driverHttp.NewLinkHandler(service)

	return LinkContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}
}

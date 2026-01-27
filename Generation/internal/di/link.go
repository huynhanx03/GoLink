package di

import (
	"go-link/common/pkg/common/cache/tinylfu"
	"go-link/common/pkg/unique"

	"go-link/generation/global"
	"go-link/generation/internal/adapters/driven/cache"
	db "go-link/generation/internal/adapters/driven/db"
	driverHttp "go-link/generation/internal/adapters/driver/http"
	"go-link/generation/internal/core/service"
	"go-link/generation/internal/infrastructure/pool"
	"go-link/generation/internal/ports"
)

type LinkContainer struct {
	Repository ports.LinkRepository
	Service    ports.LinkService
	Handler    driverHttp.LinkHandler
	CodePool   *pool.ShortCode
}

func InitLinkDependencies(clientContainer *ClientContainer) *LinkContainer {
	// Node
	node, _ := unique.NewSnowflakeNode(global.Config.SnowflakeNode, global.Time1s)

	// Pool
	pool := pool.NewShortCode(node)

	// Cache
	cache := cache.NewLink(global.Redis)

	// Repository
	repository := db.NewLinkRepository()

	// Config Cache
	localCache := tinylfu.New[string, int](tinylfu.Config{
		MaxCost: 1000,
	})

	// Service
	service := service.NewLinkService(
		repository,
		pool,
		cache,
		localCache,
		clientContainer.IdentityClient,
		clientContainer.BillingClient,
	)

	// Handler
	handler := driverHttp.NewLinkHandler(service)

	return &LinkContainer{
		Repository: repository,
		Service:    service,
		Handler:    handler,
		CodePool:   pool,
	}
}

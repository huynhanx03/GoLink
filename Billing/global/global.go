package global

import (
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/logger"
	"go-link/common/pkg/settings"

	dbEnt "go-link/billing/internal/adapters/driven/db/ent"
)

var (
	Config    settings.Config
	LoggerZap *logger.LoggerZap
	EntClient *dbEnt.EntClient
	Tinylfu   cache.LocalCache[string, any]
	Redis     cache.CacheEngine
)

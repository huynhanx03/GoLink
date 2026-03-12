package global

import (
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/database/mongodb"
	"go-link/common/pkg/logger"
	"go-link/common/pkg/settings"
)

var (
	Config    settings.Config
	LoggerZap *logger.LoggerZap
	MongoDB   *mongodb.Client
	Redis     cache.CacheEngine
)

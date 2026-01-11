package global

import (
	"go-link/common/pkg/common/cache"
	"go-link/common/pkg/database/widecolumn"
	"go-link/common/pkg/logger"
	"go-link/common/pkg/settings"
	t "go-link/common/pkg/timer"
)

var (
	Config           settings.Config
	Logger           *logger.LoggerZap
	WideColumnClient widecolumn.WideColumnClient
	Redis            cache.CacheEngine
	Time10ms         t.Timer
	Time1s           t.Timer
)

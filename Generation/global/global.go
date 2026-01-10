package global

import (
	"go-link/common/pkg/database/widecolumn"
	"go-link/common/pkg/logger"
	"go-link/common/pkg/settings"
)

var (
	Config    settings.Config
	Logger    *logger.LoggerZap
	WideColumnClient widecolumn.WideColumnClient
)

package global

import (
	"go-link/common/pkg/logger"
	"go-link/common/pkg/settings"

	"go-link/identity/internal/adapters/driven/db/ent/generate"
)

var (
	Config    settings.Config
	Logger    *logger.LoggerZap
	EntClient *generate.Client
)

package infrastructure

import (
	"go-link/billing/global"
	"go-link/common/pkg/database/redis"
)

func SetupRedis() {
	config := global.Config.Redis

	engine, err := redis.NewConnection(&config)

	if err != nil {
		global.LoggerZap.Sugar().Fatalf("Failed to connect to Redis: %v", err)
	}

	global.Redis = engine
	global.LoggerZap.Sugar().Info("Connected to Redis successfully")
}

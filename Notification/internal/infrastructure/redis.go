package infrastructure

import (
	"go-link/common/pkg/database/redis"
	"go-link/notification/global"
)

// SetupRedis initializes the Redis connection.
func SetupRedis() {
	engine, err := redis.NewConnection(&global.Config.Redis)
	if err != nil {
		global.LoggerZap.Sugar().Fatalf("Failed to connect to Redis: %v", err)
	}

	global.Redis = engine
	global.LoggerZap.Sugar().Info("Connected to Redis successfully")
}

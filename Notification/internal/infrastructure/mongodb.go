package infrastructure

import (
	"go-link/common/pkg/database/mongodb"
	"go-link/notification/global"
)

// SetupMongoDB initializes the MongoDB connection.
func SetupMongoDB() {
	client, err := mongodb.New(&global.Config.MongoDB)
	if err != nil {
		global.LoggerZap.Sugar().Fatalf("Failed to connect to MongoDB: %v", err)
	}

	global.MongoDB = client
	global.LoggerZap.Sugar().Info("Connected to MongoDB successfully")
}

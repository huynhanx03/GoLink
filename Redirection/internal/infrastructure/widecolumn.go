package infrastructure

import (
	"go-link/common/pkg/database/widecolumn"

	"go-link/redirection/global"
)

func SetupWideColumn() {
	config := global.Config.WideColumn

	client := widecolumn.NewClient(&config)

	if err := client.Connect(); err != nil {
		global.LoggerZap.Sugar().Fatalf("Failed to connect to WideColumn: %v", err)
	}

	global.WideColumnClient = client
	global.LoggerZap.Sugar().Info("Connected to WideColumn successfully")
}

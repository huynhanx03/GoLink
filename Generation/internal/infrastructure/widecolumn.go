package infrastructure

import (
	"go-link/common/pkg/database/widecolumn"

	"go-link/generation/global"
)

func SetupWideColumn() {
	config := global.Config.WideColumn

	client := widecolumn.NewClient(&config)

	if err := client.Connect(); err != nil {
		global.Logger.Sugar().Fatalf("Failed to connect to WideColumn: %v", err)
	}

	global.WideColumnClient = client
	global.Logger.Sugar().Info("Connected to WideColumn successfully")
}

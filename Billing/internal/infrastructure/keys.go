package infrastructure

import (
	"log"
	"os"

	"go-link/billing/global"
	"go-link/common/pkg/utils"
)

func SetupKeys() {
	// Load Public Key
	pubBytes, err := os.ReadFile(global.Config.JWT.PublicKeyPath)
	if err != nil {
		log.Fatalf("failed to read public key: %v", err)
	}

	global.Config.JWT.PublicKey, err = utils.ParseRSAPublicKey(pubBytes)
	if err != nil {
		log.Fatalf("failed to parse public key: %v", err)
	}
}

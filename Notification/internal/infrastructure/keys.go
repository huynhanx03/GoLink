package infrastructure

import (
	"log"
	"os"

	"go-link/common/pkg/utils"
	"go-link/notification/global"
)

// SetupKeys loads the RSA public key used to verify JWT tokens in the
// Authentication middleware. The Notification service only needs the public
// key (it never issues tokens itself).
func SetupKeys() {
	pubBytes, err := os.ReadFile(global.Config.JWT.PublicKeyPath)
	if err != nil {
		log.Fatalf("notification: failed to read public key: %v", err)
	}

	global.Config.JWT.PublicKey, err = utils.ParseRSAPublicKey(pubBytes)
	if err != nil {
		log.Fatalf("notification: failed to parse public key: %v", err)
	}
}

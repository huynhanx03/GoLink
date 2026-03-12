package main

import (
	"log"

	"go-link/notification/internal/infrastructure"
)

func main() {
	if err := infrastructure.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

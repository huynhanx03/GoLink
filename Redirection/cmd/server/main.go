package main

import (
	"go-link/redirection/internal/infrastructure"
	"log"
)

func main() {
	if err := infrastructure.Run(); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

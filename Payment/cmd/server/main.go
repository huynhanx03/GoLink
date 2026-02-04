package main

import (
	"log"

	"go-link/payment/internal/infrastructure"
)

func main() {
	if err := infrastructure.Run(); err != nil {
		log.Fatalf("Payment Service failed to run: %v", err)
	}
}

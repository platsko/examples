package main

import (
	"log"

	"lab/go-rest-api/config"
	"lab/go-rest-api/database"
	"lab/go-rest-api/router"
)

func main() {
	// Load configurations.
	if err := config.Load(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}

	// Open database connection.
	if err := database.Open(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}
	defer database.Close()

	// Start http service with routing.
	// IMPORTANT: should be called last off all,
	// because it's running in infinity loop.
	if err := router.Start(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}
}

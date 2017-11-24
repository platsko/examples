// Copyright © 2020 The EVEN Lab Team

package main

import (
	"log"

	"evenlab/go-priority-api/config"
	"evenlab/go-priority-api/database"
	"evenlab/go-priority-api/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}

	if err := database.Open(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}
	defer database.Close()

	if err := router.Start(); err != nil {
		log.Fatalf("Fatal error: %+v\n", err)
	}
}

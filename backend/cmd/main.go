package main

import (
	"log"

	"mika/internal/config"
	"mika/internal/database"
	"mika/internal/router"
)

func main() {
	cfg := config.Load()

	database.Init(cfg.DBPath)

	r := router.Setup()

	log.Printf("Server starting on :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

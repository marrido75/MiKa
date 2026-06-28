package main

import (
	"log"

	"moka/internal/router"
	"moka/internal/service"
)

func main() {
	service.InitDB()

	r := router.Setup()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

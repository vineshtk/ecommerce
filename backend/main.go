package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"
	"log"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize routes
	router := routes.SetupRoutes()

	// Start the server
	log.Println("Server running at http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

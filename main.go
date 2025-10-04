package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db := config.ConnectDatabase()

	// Migrate DB
	config.MigrateDatabase(db)

	// Init Router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Run server
	port := ":8080"
	log.Println("🚀 Server running on", port)
	r.Run(port)
}



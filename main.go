package main

import (
	"ecommerce-backend/config"
	"ecommerce-backend/routes"
	"ecommerce-backend/seeders"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	db := configs.ConnectDatabase()
	log.Println("✅ Database connected!")

	// Jalankan seeder kalau perlu
	// (sementara bisa di-comment kalau tidak mau isi data otomatis tiap start)
	seeders.Seed1(db)

	// Init Router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Ambil port dari env (Railway akan kasih PORT otomatis)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default di local
	}

	// Run server
	log.Printf("🚀 Server running on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}


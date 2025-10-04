package config

import (
	"ecommerce-backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Ambil dari environment (Railway / .env lokal)
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Kalau variabel kosong, fallback ke default (biar tetap bisa di local)
	if user == "" {
		user = "root"
	}
	if pass == "" {
		pass = ""
	}
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "3306"
	}
	if name == "" {
		name = "ecommerce"
	}

	// DSN untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB: ", err)
	}

	// Auto migrate semua tabel
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate DB: ", err)
	}

	fmt.Println("✅ Database migrated to MySQL!")
	return db
}


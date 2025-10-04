package configs

import (
	"ecommerce-backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Ambil dari ENV
	user := os.Getenv("MYSQLUSER")
	pass := os.Getenv("MYSQLPASSWORD")
	host := os.Getenv("MYSQLHOST")
	port := os.Getenv("MYSQLPORT")
	dbname := os.Getenv("MYSQLDATABASE")

	// Buat DSN MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, dbname)

	// Connect ke DB
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	// Auto migrate tabel
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate DB:", err)
	}

	fmt.Println("✅ Database migrated successfully!")
	return db
}



package config

import (
	"ecommerce-backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	// Load .env file (kalau ada)
	_ = godotenv.Load()

	// Ambil dari Railway MYSQL_URL (format mysql://user:pass@host:port/dbname)
	mysqlURL := os.Getenv("MYSQL_URL")
	if mysqlURL == "" {
		// Kalau MYSQL_URL kosong → rakit manual dari variabel
		user := os.Getenv("MYSQLUSER")
		pass := os.Getenv("MYSQLPASSWORD")
		host := os.Getenv("MYSQLHOST")
		port := os.Getenv("MYSQLPORT")
		name := os.Getenv("MYSQLDATABASE")

		mysqlURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, name,
		)
	}

	// Connect DB
	db, err := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{})
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

	fmt.Println("✅ Database connected & migrated!")
	return db
}





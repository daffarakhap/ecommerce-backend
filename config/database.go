package config

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	// Railway otomatis kasih MYSQL_URL di Variables
	dsn := os.Getenv("MYSQL_URL")
	if dsn == "" {
		log.Fatal("❌ MYSQL_URL not set in environment variables")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	log.Println("✅ Successfully connected to DB!")
	DB = db
	return db
}




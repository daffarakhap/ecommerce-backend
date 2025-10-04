package config

import (
	"log"
	"ecommerce-backend/models"
	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("❌ Failed to migrate DB: ", err)
	}

	log.Println("✅ Database migrated successfully!")
}

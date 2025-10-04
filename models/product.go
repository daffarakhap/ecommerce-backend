package models

type Product struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Stock      int      `json:"stock"`
	CategoryID uint     `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}


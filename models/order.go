package models

import "time"

type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	User      User        `json:"-" gorm:"foreignKey:UserID"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"created_at"`
}


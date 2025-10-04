package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"` // password hidden
	Role      string    `json:"role" gorm:"default:customer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Orders    []Order   `json:"-"`
}



package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}


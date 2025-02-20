package models

import (
	"time"
)

// OrderItem model, links orders with products
type OrderItem struct {
	OrderID   uint      `gorm:"primaryKey" json:"order_id"`
	ProductID uint      `gorm:"primaryKey" json:"item_id"`
	Order     Order     `gorm:"foreignKey:OrderID;references:ID"`
	Product   Product   `gorm:"foreignKey:ProductID;references:ID"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

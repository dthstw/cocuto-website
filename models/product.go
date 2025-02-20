package models

import (
	"time"
)

// Product model for the shop
type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique" json:"name"`
	Price     float64   `json:"price"`
	ImageUrl  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

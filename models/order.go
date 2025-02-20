package models

import "time"

// Order model (Tracks purchases)
type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `gorm:"index" json:"user_id"`
	User      User        `gorm:"foreignKey:UserID;references:ID"`
	Total     float64     `json:"total_price"`
	Status    string      `json:"status"` //Pending, completed, shipped
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

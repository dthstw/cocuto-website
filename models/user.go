package models

import (
	"time"

	"gorm.io/gorm"
)

// User model, tracks registered users
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"unique" json:"name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

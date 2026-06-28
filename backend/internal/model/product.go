package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Price       float64        `json:"price" gorm:"not null"`
	Category    string         `json:"category" gorm:"index"`
	ImageURL    string         `json:"image_url"`
	Stock       int            `json:"stock" gorm:"default:0"`
	Status      string         `json:"status" gorm:"default:active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

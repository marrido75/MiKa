package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description"`
	Price       float64 `json:"price" gorm:"not null"`
	Category    string  `json:"category" gorm:"index"`
	ImageURL    string  `json:"image_url"`
	Stock       int     `json:"stock" gorm:"default:0"`
	Status      string  `json:"status" gorm:"default:active"`
}

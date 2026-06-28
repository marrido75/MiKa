package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	ProductID uint           `json:"product_id" gorm:"index;not null"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	OrderID   *uint          `json:"order_id" gorm:"index"`
	Status    string         `json:"status" gorm:"default:unused"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

package model

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"index;not null"`
	Content   string `json:"content" gorm:"type:text;not null"`
	OrderID   *uint  `json:"order_id" gorm:"index"`
	Status    string `json:"status" gorm:"default:unused"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	OrderNo     string         `json:"order_no" gorm:"uniqueIndex;not null"`
	UserID      uint           `json:"user_id" gorm:"not null"`
	ProductID   uint           `json:"product_id" gorm:"not null"`
	ProductName string         `json:"product_name" gorm:"not null"`
	Quantity    int            `json:"quantity" gorm:"not null"`
	UnitPrice   float64        `json:"unit_price" gorm:"not null"`
	TotalPrice  float64        `json:"total_price" gorm:"not null"`
	CouponCode  string         `json:"coupon_code"`
	Discount    float64        `json:"discount" gorm:"default:0"`
	Status      string         `json:"status" gorm:"default:pending"`
	Cards       []Card         `json:"cards" gorm:"foreignKey:OrderID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

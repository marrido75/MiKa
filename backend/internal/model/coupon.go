package model

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Code          string         `json:"code" gorm:"uniqueIndex;not null"`
	DiscountType  string         `json:"discount_type" gorm:"not null"`
	DiscountValue float64        `json:"discount_value" gorm:"not null"`
	MinAmount     float64        `json:"min_amount" gorm:"default:0"`
	ExpiresAt     time.Time      `json:"expires_at"`
	UsageLimit    int            `json:"usage_limit" gorm:"default:0"`
	UsedCount     int            `json:"used_count" gorm:"default:0"`
	Status        string         `json:"status" gorm:"default:active"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

func NowAddDays(days int) time.Time {
	return time.Now().AddDate(0, 0, days)
}

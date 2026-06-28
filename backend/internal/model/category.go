package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"uniqueIndex;not null"`
	Label     string         `json:"label" gorm:"not null"`
	Sort      int            `json:"sort" gorm:"default:0"`
	Status    string         `json:"status" gorm:"default:active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

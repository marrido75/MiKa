package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"mika/internal/model"
)

var DB *gorm.DB

func Init(dbPath string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Card{},
		&model.Order{},
		&model.Coupon{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	SyncAllStock()
	log.Println("Database connected and migrated")
}

func SyncAllStock() {
	var products []model.Product
	DB.Find(&products)
	for _, p := range products {
		var count int64
		DB.Model(&model.Card{}).Where("product_id = ? AND status = ?", p.ID, "unused").Count(&count)
		DB.Model(&model.Product{}).Where("id = ?", p.ID).Update("stock", count)
	}
}

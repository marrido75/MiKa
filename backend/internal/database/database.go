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

	log.Println("Database connected and migrated")
}

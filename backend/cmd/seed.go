package main

import (
	"fmt"
	"log"
	"math/rand"
	"mika/internal/config"
	"mika/internal/database"
	"mika/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg := config.Load()
	database.Init(cfg.DBPath)

	// Create admin user
	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := model.User{
		Username:     "admin",
		Email:        "admin@mika.com",
		PasswordHash: string(hash),
		Role:         "admin",
	}
	database.DB.Where("username = ?", "admin").FirstOrCreate(&admin)

	// Create products
	products := []model.Product{
		{Name: "Steam 50元钱包充值码", Description: "Steam平台50元钱包充值码，即买即用", Price: 48.5, Category: "game", Stock: 10, Status: "active"},
		{Name: "Netflix 月度会员", Description: "Netflix 标准版月度会员共享账号", Price: 25.0, Category: "vip", Stock: 20, Status: "active"},
		{Name: "Spotify Premium", Description: "Spotify 高级会员3个月", Price: 35.0, Category: "vip", Stock: 15, Status: "active"},
		{Name: "JetBrains 全家桶", Description: "JetBrains 全产品年度授权码", Price: 199.0, Category: "software", Stock: 5, Status: "active"},
		{Name: "Epic 7天试用码", Description: "Epic Games 会员7天免费试用", Price: 5.0, Category: "game", Stock: 50, Status: "active"},
		{Name: "ChatGPT Plus 共享", Description: "ChatGPT Plus 账号共享，1个月", Price: 68.0, Category: "other", Stock: 8, Status: "active"},
	}

	for i := range products {
		database.DB.Where("name = ?", products[i].Name).FirstOrCreate(&products[i])
	}

	// Create cards for each product
	var dbProducts []model.Product
	database.DB.Find(&dbProducts)
	for _, p := range dbProducts {
		for j := 1; j <= p.Stock; j++ {
			card := model.Card{
				ProductID: p.ID,
				Content:   generateCode(p.Category, j),
				Status:    "unused",
			}
			database.DB.Where("content = ?", card.Content).FirstOrCreate(&card)
		}
	}

	// Create coupons
	coupon := model.Coupon{
		Code:          "WELCOME10",
		DiscountType:  "percent",
		DiscountValue: 10,
		MinAmount:     10,
		ExpiresAt:     model.NowAddDays(30),
		UsageLimit:    100,
		Status:        "active",
	}
	database.DB.Where("code = ?", "WELCOME10").FirstOrCreate(&coupon)

	coupon2 := model.Coupon{
		Code:          "SAVE5",
		DiscountType:  "fixed",
		DiscountValue: 5,
		MinAmount:     30,
		ExpiresAt:     model.NowAddDays(60),
		UsageLimit:    50,
		Status:        "active",
	}
	database.DB.Where("code = ?", "SAVE5").FirstOrCreate(&coupon2)

	log.Println("Seed data created successfully!")
}

func generateCode(category string, index int) string {
	prefix := map[string]string{"game": "GAME", "vip": "VIP", "software": "SW", "other": "KEY"}
	return fmt.Sprintf("%s-%04d-%s", prefix[category], index, randomString(8))
}

func randomString(n int) string {
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

package service

import (
	"errors"
	"fmt"
	"mika/internal/database"
	"mika/internal/model"
	"time"
)

func CreateOrder(userID, productID uint, quantity int, couponCode string) (*model.Order, error) {
	var product model.Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		return nil, errors.New("product not found")
	}
	if product.Status != "active" {
		return nil, errors.New("product is not active")
	}
	if product.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	var cards []model.Card
	if err := database.DB.Where("product_id = ? AND status = ?", productID, "unused").Limit(quantity).Find(&cards).Error; err != nil {
		return nil, err
	}
	if len(cards) < quantity {
		return nil, errors.New("insufficient cards")
	}

	totalPrice := float64(quantity) * product.Price
	discount := 0.0

	if couponCode != "" {
		coupon, err := VerifyCoupon(couponCode, totalPrice)
		if err != nil {
			return nil, fmt.Errorf("coupon error: %w", err)
		}
		discount = CalculateDiscount(coupon, totalPrice)
		if discount > totalPrice {
			discount = totalPrice
		}
		totalPrice -= discount
		coupon.UsedCount++
		database.DB.Save(coupon)
	}

	orderNo := fmt.Sprintf("MIKA%013d%04d", time.Now().UnixMilli(), userID)
	order := &model.Order{
		OrderNo:     orderNo,
		UserID:      userID,
		ProductID:   productID,
		ProductName: product.Name,
		Quantity:    quantity,
		UnitPrice:   product.Price,
		TotalPrice:  totalPrice,
		CouponCode:  couponCode,
		Discount:    discount,
		Status:      "pending",
	}

	if err := database.DB.Create(order).Error; err != nil {
		return nil, err
	}

	for i := 0; i < quantity; i++ {
		cards[i].OrderID = &order.ID
		cards[i].Status = "assigned"
		database.DB.Save(&cards[i])
	}

	product.Stock -= quantity
	database.DB.Save(&product)

	database.DB.Preload("Cards").First(order, order.ID)
	return order, nil
}

func GetUserOrders(userID uint) ([]model.Order, error) {
	var orders []model.Order
	if err := database.DB.Where("user_id = ?", userID).Preload("Cards").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderByID(id, userID uint) (*model.Order, error) {
	var order model.Order
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).Preload("Cards").First(&order).Error; err != nil {
		return nil, errors.New("order not found")
	}
	return &order, nil
}

func GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	if err := database.DB.Preload("Cards").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

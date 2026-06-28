package service

import (
	"errors"
	"mika/internal/database"
	"mika/internal/model"
	"time"

	"gorm.io/gorm"
)

func VerifyCoupon(code string, amount float64) (*model.Coupon, error) {
	var coupon model.Coupon
	if err := database.DB.Where("code = ?", code).First(&coupon).Error; err != nil {
		return nil, errors.New("coupon not found")
	}
	if coupon.Status != "active" {
		return nil, errors.New("coupon is not active")
	}
	if !coupon.ExpiresAt.IsZero() && coupon.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("coupon has expired")
	}
	if coupon.UsageLimit > 0 && coupon.UsedCount >= coupon.UsageLimit {
		return nil, errors.New("coupon usage limit reached")
	}
	if amount < coupon.MinAmount {
		return nil, errors.New("amount does not meet minimum requirement")
	}
	return &coupon, nil
}

func CalculateDiscount(coupon *model.Coupon, amount float64) float64 {
	if coupon.DiscountType == "fixed" {
		return coupon.DiscountValue
	}
	if coupon.DiscountType == "percent" {
		return amount * coupon.DiscountValue / 100
	}
	return 0
}

func IncrementCouponUsage(couponID uint) error {
	result := database.DB.Model(&model.Coupon{}).
		Where("id = ? AND (usage_limit = 0 OR used_count < usage_limit)", couponID).
		Update("used_count", gorm.Expr("used_count + 1"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("coupon usage limit reached")
	}
	return nil
}

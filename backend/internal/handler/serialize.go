package handler

import "mika/internal/model"

func serializeOrder(o model.Order) map[string]interface{} {
	cards := make([]map[string]interface{}, len(o.Cards))
	for i, c := range o.Cards {
		cards[i] = map[string]interface{}{
			"id":      c.ID,
			"content": c.Content,
			"status":  c.Status,
		}
	}
	return map[string]interface{}{
		"id":           o.ID,
		"order_no":     o.OrderNo,
		"user_id":      o.UserID,
		"product_id":   o.ProductID,
		"product_name": o.ProductName,
		"quantity":     o.Quantity,
		"unit_price":   o.UnitPrice,
		"total_price":  o.TotalPrice,
		"coupon_code":  o.CouponCode,
		"discount":     o.Discount,
		"status":       o.Status,
		"cards":        cards,
		"created_at":   o.CreatedAt,
	}
}

func serializeUser(u model.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID,
		"username":   u.Username,
		"email":      u.Email,
		"role":       u.Role,
		"created_at": u.CreatedAt,
	}
}

func serializeCard(c model.Card) map[string]interface{} {
	return map[string]interface{}{
		"id":         c.ID,
		"product_id": c.ProductID,
		"content":    c.Content,
		"order_id":   c.OrderID,
		"status":     c.Status,
		"created_at": c.CreatedAt,
	}
}

func serializeCoupon(c model.Coupon) map[string]interface{} {
	return map[string]interface{}{
		"id":             c.ID,
		"code":           c.Code,
		"discount_type":  c.DiscountType,
		"discount_value": c.DiscountValue,
		"min_amount":     c.MinAmount,
		"expires_at":     c.ExpiresAt,
		"usage_limit":    c.UsageLimit,
		"used_count":     c.UsedCount,
		"status":         c.Status,
		"created_at":     c.CreatedAt,
	}
}

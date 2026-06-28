package handler

import (
	"net/http"

	"mika/internal/service"

	"github.com/gin-gonic/gin"
)

type VerifyCouponRequest struct {
	Code   string  `json:"code" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
}

func VerifyCoupon(c *gin.Context) {
	var req VerifyCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon, err := service.VerifyCoupon(req.Code, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	discount := service.CalculateDiscount(coupon, req.Amount)
	finalPrice := req.Amount - discount

	c.JSON(http.StatusOK, gin.H{
		"discount":    discount,
		"final_price": finalPrice,
		"coupon":      serializeCoupon(*coupon),
	})
}

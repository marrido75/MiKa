package handler

import (
	"net/http"
	"strconv"

	"mika/internal/database"
	"mika/internal/model"
	"mika/internal/service"

	"github.com/gin-gonic/gin"
)

func AdminCreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func AdminUpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	existing, err := service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err := c.ShouldBindJSON(existing); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateProduct(existing); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, existing)
}

func AdminDeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	if err := service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

type ImportCardsRequest struct {
	Contents []string `json:"contents" binding:"required,min=1"`
}

func AdminImportCards(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req ImportCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := service.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	for _, content := range req.Contents {
		card := model.Card{
			ProductID: uint(productID),
			Content:   content,
			Status:    "unused",
		}
		database.DB.Create(&card)
	}

	existing.Stock += len(req.Contents)
	service.UpdateProduct(existing)

	c.JSON(http.StatusOK, gin.H{"message": "imported", "count": len(req.Contents)})
}

func AdminGetOrders(c *gin.Context) {
	orders, err := service.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

type CreateCouponRequest struct {
	Code          string  `json:"code" binding:"required"`
	DiscountType  string  `json:"discount_type" binding:"required"`
	DiscountValue float64 `json:"discount_value" binding:"required"`
	MinAmount     float64 `json:"min_amount"`
	ExpiresAt     string  `json:"expires_at"`
	UsageLimit    int     `json:"usage_limit"`
}

func AdminCreateCoupon(c *gin.Context) {
	var req CreateCouponRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupon := model.Coupon{
		Code:          req.Code,
		DiscountType:  req.DiscountType,
		DiscountValue: req.DiscountValue,
		MinAmount:     req.MinAmount,
		UsageLimit:    req.UsageLimit,
		Status:        "active",
	}

	if err := database.DB.Create(&coupon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func AdminGetCoupons(c *gin.Context) {
	var coupons []model.Coupon
	if err := database.DB.Find(&coupons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"coupons": coupons})
}

func AdminUpdateCoupon(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon id"})
		return
	}

	var coupon model.Coupon
	if err := database.DB.First(&coupon, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&coupon).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupon)
}

func AdminDeleteCoupon(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid coupon id"})
		return
	}

	if err := database.DB.Delete(&model.Coupon{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

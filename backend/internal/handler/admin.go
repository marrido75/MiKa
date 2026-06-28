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
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          product.ID,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"category":    product.Category,
		"image_url":   product.ImageURL,
		"stock":       product.Stock,
		"status":      product.Status,
		"created_at":  product.CreatedAt,
	})
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
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          existing.ID,
		"name":        existing.Name,
		"description": existing.Description,
		"price":       existing.Price,
		"category":    existing.Category,
		"image_url":   existing.ImageURL,
		"stock":       existing.Stock,
		"status":      existing.Status,
		"created_at":  existing.CreatedAt,
	})
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
	ProductID uint     `json:"product_id" binding:"required"`
	Contents  []string `json:"contents" binding:"required,min=1"`
}

func AdminImportCards(c *gin.Context) {
	var req ImportCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := service.GetProductByID(req.ProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	for _, content := range req.Contents {
		card := model.Card{
			ProductID: req.ProductID,
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
	list := make([]map[string]interface{}, len(orders))
	for i, o := range orders {
		list[i] = serializeOrder(o)
	}
	c.JSON(http.StatusOK, list)
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
	c.JSON(http.StatusOK, serializeCoupon(coupon))
}

func AdminGetCoupons(c *gin.Context) {
	var coupons []model.Coupon
	if err := database.DB.Find(&coupons).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	list := make([]map[string]interface{}, len(coupons))
	for i, c := range coupons {
		list[i] = serializeCoupon(c)
	}
	c.JSON(http.StatusOK, list)
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
	c.JSON(http.StatusOK, serializeCoupon(coupon))
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

func AdminGetStats(c *gin.Context) {
	var totalOrders int64
	var totalRevenue float64
	var totalProducts int64
	var totalUsers int64

	database.DB.Model(&model.Order{}).Where("status = ?", "paid").Count(&totalOrders)
	database.DB.Model(&model.Order{}).Where("status = ?", "paid").Select("COALESCE(SUM(total_price), 0)").Scan(&totalRevenue)
	database.DB.Model(&model.Product{}).Count(&totalProducts)
	database.DB.Model(&model.User{}).Count(&totalUsers)

	c.JSON(http.StatusOK, gin.H{
		"totalOrders":   totalOrders,
		"totalRevenue":  totalRevenue,
		"totalProducts": totalProducts,
		"totalUsers":    totalUsers,
	})
}

func AdminGetUsers(c *gin.Context) {
	var users []model.User
	database.DB.Select("id, username, email, role, created_at").Find(&users)
	list := make([]map[string]interface{}, len(users))
	for i, u := range users {
		list[i] = serializeUser(u)
	}
	c.JSON(http.StatusOK, list)
}

func AdminGetProductCards(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	var cards []model.Card
	database.DB.Where("product_id = ?", id).Order("created_at DESC").Find(&cards)
	list := make([]map[string]interface{}, len(cards))
	for i, card := range cards {
		list[i] = serializeCard(card)
	}
	c.JSON(http.StatusOK, list)
}

func AdminDeleteCard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid card id"})
		return
	}
	database.DB.Delete(&model.Card{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

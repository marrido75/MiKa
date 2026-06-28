package handler

import (
	"net/http"
	"strconv"

	"mika/internal/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	category := c.Query("category")
	products, total := service.GetProducts(category)
	list := make([]map[string]interface{}, len(products))
	for i, p := range products {
		list[i] = map[string]interface{}{
			"id":          p.ID,
			"name":        p.Name,
			"description": p.Description,
			"price":       p.Price,
			"category":    p.Category,
			"image_url":   p.ImageURL,
			"stock":       p.Stock,
			"status":      p.Status,
			"created_at":  p.CreatedAt,
		}
	}
	c.JSON(http.StatusOK, gin.H{"products": list, "total": total})
}

func GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
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

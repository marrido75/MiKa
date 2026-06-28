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
	c.JSON(http.StatusOK, gin.H{"products": products, "total": total})
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
	c.JSON(http.StatusOK, product)
}

package handler

import (
	"net/http"
	"strconv"

	"mika/internal/database"
	"mika/internal/model"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []model.Category
	database.DB.Where("status = ?", "active").Order("sort ASC, id ASC").Find(&categories)
	list := make([]map[string]interface{}, len(categories))
	for i, cat := range categories {
		list[i] = map[string]interface{}{
			"id":    cat.ID,
			"name":  cat.Name,
			"label": cat.Label,
			"sort":  cat.Sort,
		}
	}
	c.JSON(http.StatusOK, list)
}

func AdminGetCategories(c *gin.Context) {
	var categories []model.Category
	database.DB.Order("sort ASC, id ASC").Find(&categories)
	list := make([]map[string]interface{}, len(categories))
	for i, cat := range categories {
		list[i] = map[string]interface{}{
			"id":      cat.ID,
			"name":    cat.Name,
			"label":   cat.Label,
			"sort":    cat.Sort,
			"status":  cat.Status,
			"created_at": cat.CreatedAt,
		}
	}
	c.JSON(http.StatusOK, list)
}

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Label string `json:"label" binding:"required"`
	Sort  int    `json:"sort"`
}

func AdminCreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat := model.Category{
		Name:   req.Name,
		Label:  req.Label,
		Sort:   req.Sort,
		Status: "active",
	}

	if err := database.DB.Create(&cat).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category name already exists"})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     cat.ID,
		"name":   cat.Name,
		"label":  cat.Label,
		"sort":   cat.Sort,
		"status": cat.Status,
	})
}

func AdminUpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var cat model.Category
	if err := database.DB.First(&cat, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat.Name = req.Name
	cat.Label = req.Label
	cat.Sort = req.Sort
	database.DB.Save(&cat)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     cat.ID,
		"name":   cat.Name,
		"label":  cat.Label,
		"sort":   cat.Sort,
		"status": cat.Status,
	})
}

func AdminDeleteCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	database.DB.Delete(&model.Category{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Go + Gin backend!",
		"time":    "2026",
	})
}

func GetUsers(c *gin.Context) {
	users := []gin.H{
		{"id": 1, "name": "Alice", "email": "alice@example.com"},
		{"id": 2, "name": "Bob", "email": "bob@example.com"},
		{"id": 3, "name": "Charlie", "email": "charlie@example.com"},
	}
	c.JSON(http.StatusOK, users)
}

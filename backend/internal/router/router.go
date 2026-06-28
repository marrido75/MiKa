package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"mika/internal/handler"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api/health", handler.Health)
	r.GET("/api/message", handler.GetMessage)
	r.GET("/api/users", handler.GetUsers)

	return r
}

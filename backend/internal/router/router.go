package router

import (
	"mika/internal/config"
	"mika/internal/handler"
	"mika/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(cfg *config.Config) *gin.Engine {
	middleware.SetSecret(cfg.JWTSecret)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/auth/register", handler.Register)
		api.POST("/auth/login", handler.Login)
		api.GET("/products", handler.GetProducts)
		api.GET("/products/:id", handler.GetProduct)
		api.POST("/coupons/verify", handler.VerifyCoupon)

		auth := api.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("/auth/profile", handler.GetProfile)
			auth.POST("/orders", handler.CreateOrder)
			auth.GET("/orders/:id", handler.GetOrder)
			auth.GET("/user/orders", handler.GetUserOrders)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired(), middleware.AdminRequired())
		{
			admin.POST("/products", handler.AdminCreateProduct)
			admin.PUT("/products/:id", handler.AdminUpdateProduct)
			admin.DELETE("/products/:id", handler.AdminDeleteProduct)
			admin.POST("/cards/import", handler.AdminImportCards)
			admin.GET("/orders", handler.AdminGetOrders)
			admin.POST("/coupons", handler.AdminCreateCoupon)
			admin.GET("/coupons", handler.AdminGetCoupons)
			admin.PUT("/coupons/:id", handler.AdminUpdateCoupon)
			admin.DELETE("/coupons/:id", handler.AdminDeleteCoupon)
		}
	}

	return r
}

package routes

import (
	"pos-app/handlers"
	"pos-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", handlers.Register)
			auth.POST("/login", handlers.Login)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			products := protected.Group("/products")
			{
				products.GET("", handlers.GetProducts)
				products.POST("", middleware.AdminOnly(), handlers.CreateProduct)
				products.PUT("/:id", middleware.AdminOnly(), handlers.UpdateProduct)
				products.DELETE("/:id", middleware.AdminOnly(), handlers.DeleteProduct)
			}
		}
	}
}

package routes

import (
	middleware "backend_go/middlewares"
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.Use(middleware.AuthMiddleware())
		products.GET("/", services.GetProducts)
		products.GET("/:id", services.GetProductByID)
		products.POST("/", services.CreateProduct)
		products.PATCH("/:id", services.UpdateProduct)
		products.DELETE("/:id", services.DeleteProduct)
	}
}

package routes

import (
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	r.GET("/products", services.GetAllProducts)
	r.GET("/products/:id", services.GetProductByID)
}

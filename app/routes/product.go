package routes

import (
	middleware "backend_go/app/middlewares"
	"backend_go/app/permissions"
	"backend_go/app/services"
	"backend_go/constants"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.GET("/", services.GetProducts)
		products.GET("/:id", services.GetProductByID)
	}

	protected := r.Group("/products").Use(
		middleware.AuthMiddlewareFlexible(),
	)
	{
		protected.POST("/", permissions.RequirePermissions(constants.PermissionProductsCreate), services.CreateProduct)
		protected.PATCH("/:id",
			permissions.RequirePermissions(constants.PermissionProductsUpdate),
			services.UpdateProduct,
		)
		protected.DELETE("/:id",
			permissions.RequirePermissions(constants.PermissionProductsDelete),
			services.DeleteProduct,
		)
	}
}

package routes

import (
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/", services.Register)
		auth.POST("/", services.Login)
	}
}

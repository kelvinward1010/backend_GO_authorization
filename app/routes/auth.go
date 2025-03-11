package routes

import (
	"backend_go/app/services"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", services.Register)
		auth.POST("/login", services.Login)
	}
}

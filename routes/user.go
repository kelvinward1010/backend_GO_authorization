package routes

import (
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/user")
	{
		users.GET("/", services.GetUsers)
		users.GET("/:id", services.GetUserByID)
		users.PATCH("/:id", services.UpdateUser)
		users.DELETE("/:id", services.DeleteUser)
	}
}

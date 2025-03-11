package routes

import (
	"backend_go/constants"
	middleware "backend_go/middlewares"
	"backend_go/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	users := r.Group("/user")
	{
		users.GET("/", middleware.RequireRoles(constants.RoleAdmin), services.GetUsers)
		users.GET("/:id", middleware.RequireRoles(constants.RoleAdmin), services.GetUserByID)
		users.PATCH("/:id", middleware.RequireRoles(constants.RoleAdmin), services.UpdateUser)
		users.DELETE("/:id", middleware.RequireRoles(constants.RoleAdmin), services.DeleteUser)
	}
}

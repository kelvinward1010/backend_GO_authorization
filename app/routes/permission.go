package routes

import (
	middleware "backend_go/app/middlewares"
	"backend_go/app/permissions"
	"backend_go/app/services"
	"backend_go/constants"

	"github.com/gin-gonic/gin"
)

func PermissionRoutes(r *gin.RouterGroup) {
	permissionsGroup := r.Group("/permissions")

	permissionsGroup.GET("/",
		middleware.AuthMiddlewareFlexible(),
		permissions.RequirePermissions(constants.PermissionRolesGet),
		services.GetRoles,
	)

	permissionsGroup.GET("/all",
		middleware.AuthMiddlewareFlexible(),
		permissions.RequirePermissions(constants.PermissionRolesGet),
		services.GetAllPermissions,
	)

	permissionsGroup.PATCH("/",
		middleware.AuthMiddlewareFlexible(),
		permissions.RequirePermissions(constants.PermissionRolesUpdate),
		services.UpdateRolePermissions,
	)

	permissionsGroup.PATCH("/users/:id",
		middleware.AuthMiddlewareFlexible(),
		permissions.RequirePermissions(constants.PermissionUsersUpdate),
		services.UpdateUserPermissions,
	)

	permissionsGroup.DELETE("/:id",
		middleware.AuthMiddlewareFlexible(),
		permissions.RequirePermissions(constants.PermissionRolesDelete),
		services.DeleteRole,
	)
}

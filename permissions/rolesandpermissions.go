package permissions

import (
	"backend_go/constants"
	"backend_go/models"
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SeedRolesAndPermissions(db *gorm.DB) {
	var roles = []models.Role{
		{Name: constants.RoleAdmin},
		{Name: constants.RoleSales},
		{Name: constants.RoleUser},
	}

	var permissions = []models.Permission{
		{Name: constants.PermissionUsersGet}, {Name: constants.PermissionUsersCreate}, {Name: constants.PermissionUsersUpdate}, {Name: constants.PermissionUsersDelete},
		{Name: constants.PermissionProductsGet}, {Name: constants.PermissionProductsCreate}, {Name: constants.PermissionProductsUpdate}, {Name: constants.PermissionProductsDelete},
		{Name: constants.PermissionRolesGet}, {Name: constants.PermissionRolesUpdate}, {Name: constants.PermissionRolesDelete},
	}

	for i := range permissions {
		db.FirstOrCreate(&permissions[i], models.Permission{Name: permissions[i].Name})
	}

	for i := range roles {
		db.FirstOrCreate(&roles[i], models.Role{Name: roles[i].Name})
	}

	var allPermissions []models.Permission
	db.Find(&allPermissions)

	var productPermissions []models.Permission
	db.Where("name LIKE ?", "products:%").Find(&productPermissions)

	var userPermissions []models.Permission
	db.Where("name = ?", constants.PermissionProductsGet).Find(&userPermissions)

	var rolePermissions []models.Permission
	db.Where("name = ?", constants.PermissionProductsGet).Find(&rolePermissions)

	db.Model(&roles[0]).Association("Permissions").Replace(allPermissions)
	db.Model(&roles[1]).Association("Permissions").Replace(productPermissions)
	db.Model(&roles[2]).Association("Permissions").Replace(userPermissions)
	db.Model(&roles[2]).Association("Permissions").Replace(rolePermissions)
}

func RequirePermissions(allowedPermissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, exists := c.Get("permissions")
		if !exists {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing permissions")
			c.Abort()
			return
		}

		userPermissions, ok := permissions.([]string)
		if !ok {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Internal error: Invalid permissions format")
			c.Abort()
			return
		}

		for _, allowed := range allowedPermissions {
			for _, userPerm := range userPermissions {
				if userPerm == allowed {
					c.Next()
					return
				}
			}
		}

		utils.SendErrorResponse(c, http.StatusForbidden, "Forbidden: Insufficient permissions")
		c.Abort()
	}
}

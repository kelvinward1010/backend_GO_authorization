package middleware

import (
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesIface, exists := c.Get("roles")
		if !exists {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing roles")
			c.Abort()
			return
		}

		roles, ok := rolesIface.([]string)
		if !ok {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Internal error: Invalid roles format")
			c.Abort()
			return
		}

		for _, userRole := range roles {
			for _, allowed := range allowedRoles {
				if userRole == allowed {
					c.Next()
					return
				}
			}
		}

		utils.SendErrorResponse(c, http.StatusForbidden, "Forbidden: Insufficient permissions")
		c.Abort()
	}
}

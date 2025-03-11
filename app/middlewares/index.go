package middleware

import (
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing role")
			c.Abort()
			return
		}

		for _, allowed := range allowedRoles {
			if role == allowed {
				c.Next()
				return
			}
		}

		utils.SendErrorResponse(c, http.StatusForbidden, "Forbidden: Insufficient permissions")
		c.Abort()
	}
}

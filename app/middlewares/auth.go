package middleware

import (
	"backend_go/core"
	"backend_go/models"
	"backend_go/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Token missing or invalid")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, claims, err := core.VerifyToken(tokenString)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid token payload")
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing role information")
			c.Abort()
			return
		}

		c.Set("user_id", int(userID))
		c.Set("role", role)

		c.Next()
	}
}

func AuthMiddlewareFlexible() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Token missing or invalid")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		_, claims, err := core.VerifyToken(tokenString)
		if err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid token")
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Invalid token payload")
			c.Abort()
			return
		}

		roleName, ok := claims["role"].(string)
		if !ok {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: Missing role information")
			c.Abort()
			return
		}

		var role models.Role
		if err := core.DB.Preload("Permissions").Where("name = ?", roleName).First(&role).Error; err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Error fetching role permissions")
			c.Abort()
			return
		}

		userPermissions := make([]string, len(role.Permissions))
		for i, perm := range role.Permissions {
			userPermissions[i] = perm.Name
		}

		c.Set("user_id", int(userID))
		c.Set("role", roleName)
		c.Set("permissions", userPermissions)

		c.Next()
	}
}

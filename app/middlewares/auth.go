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

		var user models.User
		if err := core.DB.Preload("Roles.Permissions").Preload("Permissions").First(&user, uint(userID)).Error; err != nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Unauthorized: User not found")
			c.Abort()
			return
		}

		permMap := make(map[string]struct{})
		for _, role := range user.Roles {
			for _, perm := range role.Permissions {
				permMap[perm.Name] = struct{}{}
			}
		}
		for _, perm := range user.Permissions {
			permMap[perm.Name] = struct{}{}
		}

		var allPerms []string
		for name := range permMap {
			allPerms = append(allPerms, name)
		}

		var roleNames []string
		for _, r := range user.Roles {
			roleNames = append(roleNames, r.Name)
		}

		c.Set("user_id", uint(userID))
		c.Set("roles", roleNames)
		c.Set("permissions", allPerms)

		c.Next()
	}
}

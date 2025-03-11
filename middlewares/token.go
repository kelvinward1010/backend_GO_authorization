package middleware

import (
	"backend_go/core"
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

		c.Set("user_id", int(userID))
		c.Next()
	}
}

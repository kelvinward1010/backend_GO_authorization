package app

import (
	"backend_go/app/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin API"})
	})

	api := r.Group("/api/v1")
	routes.ProductRoutes(api)
	routes.AuthRoutes(api)
	routes.UserRoutes(api)

	return r
}

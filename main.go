package main

import (
	routes "backend_go/app"
	_ "backend_go/app/docs"
	db "backend_go/core"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Multi-Feature API Documentation
// @version 1.0
// @description API Documentation for Users and Products
// @host localhost:8000
// @BasePath /api/v1
func main() {
	db.ConnectDB()

	DBConn, err := db.DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database connection: %v", err)
	}
	defer DBConn.Close()

	r := routes.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("🚀 Server running on port 8000")
	r.Run(":8000")
}

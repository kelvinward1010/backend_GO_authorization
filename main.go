package main

import (
	db "backend_go/core"
	"backend_go/routes"
	"fmt"
	"log"
)

func main() {
	db.ConnectDB()

	DBConn, err := db.DB.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get database connection: %v", err)
	}
	defer DBConn.Close()

	r := routes.SetupRouter()

	fmt.Println("🚀 Server running on port 8000")
	r.Run(":8000")
}

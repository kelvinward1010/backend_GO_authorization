package main

import (
	"backend_go/routes"
)

func main() {
	r := routes.SetupRouter()
	routes.SetupRouter()

	r.Run(":8000")
}

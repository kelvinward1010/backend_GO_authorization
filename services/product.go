package services

import (
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	products := []string{"Product A", "Product B", "Product C"}

	utils.SendResponse(c, http.StatusOK, "success", products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product := map[string]string{"id": id, "name": "Product " + id}

	utils.SendResponse(c, http.StatusOK, "Get product by ID", product)
}

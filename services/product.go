package services

import (
	"backend_go/db"
	"backend_go/models"
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)

	utils.SendResponse(c, http.StatusOK, "Product list retrieved successfully", products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Product details retrieved successfully", product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Create(&product).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.SendResponse(c, http.StatusCreated, "Product created successfully", product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	db.DB.Save(&product)
	utils.SendResponse(c, http.StatusOK, "Product updated successfully", product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	db.DB.Delete(&product)
	utils.SendResponse(c, http.StatusOK, "Product deleted successfully", nil)
}

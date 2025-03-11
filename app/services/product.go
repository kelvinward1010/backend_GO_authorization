package services

import (
	db "backend_go/core"
	"backend_go/models"
	"backend_go/models/schemas"
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetProducts godoc
// @Summary Get all products
// @Tags Products
// @Description Retrieve list of products with optional name filter
// @Produce json
// @Param name query string false "Filter by product name (case-insensitive)"
// @Success 200 {object} utils.ResponseGin{data=[]models.Product}
// @Failure 500 {object} utils.ResponseGin
// @Router /products [get]
func GetProducts(c *gin.Context) {
	name := c.Query("name")

	var products []models.Product
	query := db.DB

	if name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	result := query.Find(&products)
	if result.Error != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Products retrieved successfully", products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Tags Products
// @Description Get product details by ID
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.ResponseGin{data=models.Product}
// @Failure 404 {object} utils.ResponseGin
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := db.DB.First(&product, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Product details retrieved successfully", product)
}

// CreateProduct godoc
// @Summary Create new product
// @Tags Products
// @Description Create a new product
// @Accept json
// @Produce json
// @Param product body schemas.ProductCreateRequest true "Product Data"
// @Success 201 {object} utils.ResponseGin{data=models.Product}
// @Failure 400 {object} utils.ResponseGin
// @Failure 500 {object} utils.ResponseGin
// @Router /products [post]
// @Security BearerAuth
func CreateProduct(c *gin.Context) {
	var product schemas.ProductCreateRequest

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newProduct := models.Product{
		Name:  product.Name,
		Price: product.Price,
	}

	if err := db.DB.Create(&newProduct).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.SendResponse(c, http.StatusCreated, "Product created successfully", product)
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Products
// @Description Update existing product
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body schemas.ProductUpdateRequest true "Updated product data"
// @Success 200 {object} utils.ResponseGin{data=models.Product}
// @Failure 400 {object} utils.ResponseGin
// @Failure 404 {object} utils.ResponseGin
// @Router /products/{id} [patch]
// @Security BearerAuth
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	var updateData schemas.ProductUpdateRequest

	if err := db.DB.First(&product, id).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Product not found")
		return
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updates := map[string]interface{}{
		"Name":  updateData.Name,
		"Price": updateData.Price,
	}

	if err := db.DB.Model(&product).Updates(updates).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description Delete product by ID
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} utils.ResponseGin
// @Failure 404 {object} utils.ResponseGin
// @Router /products/{id} [delete]
// @Security BearerAuth
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

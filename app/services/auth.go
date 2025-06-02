package services

import (
	"backend_go/core"
	"backend_go/models"
	"backend_go/models/schemas"
	"backend_go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	hashedPassword, err := core.HashPassword(user.Password)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}
	user.Password = hashedPassword

	if err := core.DB.Create(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to register user")
		return
	}

	utils.SendResponse(c, http.StatusCreated, "User registered successfully", nil)
}

func Login(c *gin.Context) {
	var req schemas.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	var user models.User
	if err := core.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	if !core.CheckPassword(user.Password, req.Password) {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid password")
		return
	}

	permissionNames := make([]string, len(user.Permissions))
	for i, perm := range user.Permissions {
		permissionNames[i] = perm.Name
	}
	token, err := core.GenerateTokenWithPermissions(int(user.ID), user.Username, user.Role, permissionNames)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	user.Password = ""

	utils.SendResponse(c, http.StatusOK, "Login successful", gin.H{
		"user":  user,
		"token": token,
	})
}

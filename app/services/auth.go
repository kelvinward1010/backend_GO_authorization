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
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	hashedPassword, err := core.HashPassword(input.Password)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	var roles []models.Role
	roleIDs := []uint{}
	for _, r := range input.Roles {
		roleIDs = append(roleIDs, r.ID)
	}
	if len(roleIDs) > 0 {
		if err := core.DB.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
			utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to fetch roles")
			return
		}
		user.Roles = roles
	}

	if err := core.DB.Create(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
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

	roleNames := make([]string, len(user.Roles))
	for i, role := range user.Roles {
		roleNames[i] = role.Name
	}

	token, err := core.GenerateTokenWithPermissions(int(user.ID), user.Username, roleNames, permissionNames)
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

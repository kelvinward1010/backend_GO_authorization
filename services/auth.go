package services

import (
	"backend_go/core"
	"backend_go/models"
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
	var userInput models.User
	var user models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := core.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !core.CheckPassword(user.Password, userInput.Password) {
		utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := core.GenerateToken(user.Username)
	if err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Login successful", gin.H{"token": token})
}

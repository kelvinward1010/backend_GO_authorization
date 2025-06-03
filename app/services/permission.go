package services

import (
	"backend_go/core"
	"backend_go/models"
	"backend_go/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role

	if err := core.DB.Preload("Permissions").Find(&roles).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Cannot get list of roles")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Danh sách roles", roles)
}

func GetAllPermissions(c *gin.Context) {
	var permissions []models.Permission
	if err := core.DB.Find(&permissions).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to load permissions")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Permissions fetched", permissions)
}

func UpdateRolePermissions(c *gin.Context) {
	var request struct {
		RoleID      int      `json:"role_id"`
		Permissions []string `json:"permissions"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var role models.Role
	if err := core.DB.First(&role, request.RoleID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Role not found")
		return
	}

	var permissions []models.Permission
	if err := core.DB.Where("name IN ?", request.Permissions).Find(&permissions).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Error fetching permissions")
		return
	}

	if err := core.DB.Model(&role).Association("Permissions").Replace(permissions); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update role permissions")
		return
	}

	if err := core.DB.Preload("Permissions").First(&role, role.ID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to reload updated role")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Role permissions updated successfully", role)
}

func UpdateUserPermissions(c *gin.Context) {
	var req struct {
		Permissions []string `json:"permissions"`
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	var user models.User
	if err := core.DB.First(&user, userID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	var permissions []models.Permission
	if err := core.DB.Where("name IN ?", req.Permissions).Find(&permissions).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Error loading permissions")
		return
	}

	if err := core.DB.Model(&user).Association("Permissions").Replace(permissions); err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update user permissions")
		return
	}

	if err := core.DB.Preload("Permissions").First(&user, user.ID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to reload updated user")
		return
	}

	utils.SendResponse(c, http.StatusOK, "User permissions updated", user)
}

func DeleteRole(c *gin.Context) {
	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var role models.Role
	if err := core.DB.First(&role, roleID).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusNotFound, "Role does not exist")
		return
	}

	if err := core.DB.Select("Permissions").Delete(&role).Error; err != nil {
		utils.SendErrorResponse(c, http.StatusInternalServerError, "Cannot delete role")
		return
	}

	utils.SendResponse(c, http.StatusOK, "Successfully deleted role", nil)
}

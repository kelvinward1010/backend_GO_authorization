package utils

import (
	"github.com/gin-gonic/gin"
)

type ResponseGin struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, ResponseGin{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, ResponseGin{
		Status:  status,
		Message: message,
	})
}

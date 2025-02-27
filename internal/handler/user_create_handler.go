package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var userModel *dto.CreateUserDTO

	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(500, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.CreateUser(userModel)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user is created successfully"})
}

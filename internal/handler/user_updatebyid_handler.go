package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser *dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(500, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.UpdateUserByID(id, updatedUser)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "user is updated successfully"})
}

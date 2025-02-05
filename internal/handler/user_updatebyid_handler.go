package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}

	var updatedUser *entity.UpdateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.s.UpdateUserByID(id, updatedUser)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// OK
	c.JSON(200, gin.H{"message": "User updated successfully"})
}

package handler

import (
	"firstGoProject/internal/domain/entity"
	"fmt"
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

	var updatedUser entity.User

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	updatedUser.ID = id

	err = h.service.UpdateUserByID(id, updatedUser)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	// OK
	c.JSON(200, gin.H{"message": "User updated successfully"})
}

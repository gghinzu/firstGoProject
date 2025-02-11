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
		c.JSON(400, gin.H{"error": "id should be numeric"})
		return
	}

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser *entity.UpdateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(500, gin.H{"error": "json cannot be bound"})
		return
	}

	err = h.s.UpdateUserByID(id, updatedUser)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
}

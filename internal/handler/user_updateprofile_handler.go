package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var updateData *dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := strings.TrimPrefix(c.Param("id"), "/")

	_, err := h.s.UpdateProfile(id, updateData)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "profile is updated successfully"})
}

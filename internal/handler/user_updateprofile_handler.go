package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var updateData *dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := strings.TrimPrefix(c.Param("id"), "/")

	_, err := h.s.UpdateProfile(id, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile is updated successfully"})
}

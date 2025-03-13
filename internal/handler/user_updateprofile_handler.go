package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var updateData *dto.UpdateProfileDTO
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims := c.MustGet("claims").(helper.UserCustomClaims)

	_, err := h.s.UpdateProfile(claims.ID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile is updated successfully"})
}

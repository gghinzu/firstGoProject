package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser *dto.UpdateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.UpdateUserByID(id, updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user is updated successfully"})
}

package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var info *dto.SignUpDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.s.SignUp(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user signed up successfully"})
}

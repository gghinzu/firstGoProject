package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var info *dto.SignUpDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.s.SignUp(info)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user signed up successfully"})
}

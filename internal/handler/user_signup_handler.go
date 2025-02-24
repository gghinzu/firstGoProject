package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var info entity.SignUpDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON data"})
		return
	}

	err := h.s.SignUp(&info)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User signed up successfully"})
}

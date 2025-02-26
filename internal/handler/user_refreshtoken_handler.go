package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var email entity.RefreshTokenDTO
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if email.Email == nil {
		c.JSON(400, gin.H{"error": "email must be provided"})
		return
	}

	refresh, errToken := h.s.RefreshToken(email.Email)
	if errToken != nil {
		c.JSON(401, gin.H{"error": errToken.Error()})
		return
	}

	c.JSON(200, gin.H{"refresh_token": refresh.RefreshToken})

}

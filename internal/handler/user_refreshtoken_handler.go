package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var token dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := helper.VerifyToken(token.Token)
	if err != nil {
		c.JSON(400, gin.H{"error": "token is invalid"})
		return
	}

	if token.Token == "" {
		c.JSON(400, gin.H{"error": "token must be provided"})
		return
	}

	refresh, errToken := h.s.RefreshToken()
	if errToken != nil {
		c.JSON(401, gin.H{"error": errToken.Error()})
		return
	}

	c.JSON(200, gin.H{"refresh_token": refresh.RefreshToken})
}

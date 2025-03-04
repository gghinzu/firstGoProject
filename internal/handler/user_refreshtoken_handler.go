package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var refreshToken *dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := helper.VerifyToken(refreshToken.RefreshToken, "refresh")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := claims.UserID
	tokenUser, err := h.s.RefreshToken(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokenUser)

}

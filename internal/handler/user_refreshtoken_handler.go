package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var refreshToken *dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	claims, err := helper.VerifyToken(refreshToken.RefreshToken, "refresh")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	tokenUser, err := h.s.RefreshToken(claims.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": error.NotAuthorized})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access-token:": tokenUser.RefreshToken})

}

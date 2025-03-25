package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var refreshTokenDTO *dto.RefreshTokenDTO

	if err := c.ShouldBindJSON(&refreshTokenDTO); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest)
		return
	}

	claims, err := helper.VerifyToken(refreshTokenDTO.RefreshToken, "refresh")
	if err != nil {
		c.JSON(http.StatusBadRequest, error.Unauthenticated)
		return
	}

	tokenUser, err := h.s.RefreshToken(claims.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, error.Unauthorized)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenUser.Token,
	})
}

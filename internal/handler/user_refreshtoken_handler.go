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
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	claims, err := helper.VerifyToken(refreshToken.RefreshToken, "refresh")
	if err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	tokenUser, err := h.s.RefreshToken(claims.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, error.Unauthorized.Error())
		return
	}

	c.JSON(http.StatusOK, tokenUser.RefreshToken)
}

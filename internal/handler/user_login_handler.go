package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var loginInfo dto.LoginDTO

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginInfo.Password == "" || loginInfo.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password cannot be empty"})
		return
	}

	userLogin, err := h.s.Login(loginInfo)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  userLogin.Token,
		"refresh_token": userLogin.RefreshToken,
	})
}

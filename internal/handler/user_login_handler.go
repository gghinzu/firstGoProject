package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var loginInfo dto.LoginDTO

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if loginInfo.Password == "" {
		c.JSON(400, gin.H{"error": "password cannot be empty"})
		return
	}

	user, errToken := h.s.Login(loginInfo)
	if errToken != nil {
		c.JSON(401, gin.H{"error": errToken.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": user.Token,
	})
}

package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var loginInfo entity.LoginDTO

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if loginInfo.Email == nil || loginInfo.Password == nil {
		c.JSON(400, gin.H{"error": "email or password is nil"})
		return
	}

	user, errToken := h.s.Login(loginInfo)
	if errToken != nil {
		c.JSON(401, gin.H{"error": errToken.Error()})
		return
	}

	c.JSON(200, gin.H{
		"email": user.Email,
		"token": user.Token,
		/*		"refresh_token": user.RefreshToken,
		 */})
}

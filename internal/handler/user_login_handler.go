package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var loginInfo dto.LoginDTO

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	err := validate.Struct(&loginInfo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error.InvalidInput})
		return
	}

	var userLogin *dto.TokenDTO

	userLogin, err = h.s.Login(loginInfo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": error.NotAuthenticated})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  userLogin.Token,
		"refresh_token": userLogin.RefreshToken,
	})
}

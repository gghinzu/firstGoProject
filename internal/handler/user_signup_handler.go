package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var info dto.SignUpDTO

	validate := validator.New()

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email must be valid and password must be at least 6 characters"})
		return
	}

	err = h.s.SignUp(&info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user signed up successfully"})
}

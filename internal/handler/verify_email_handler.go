package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func init() {
	validate = validator.New()
}

func (h *UserHandler) VerifyEmailHandler(c *gin.Context) {
	var input dto.VerifyEmailDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	if err := validate.Struct(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error.InvalidInput})
		return
	}

	err := h.s.VerifyEmail(input.Email, input.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

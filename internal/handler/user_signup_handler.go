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

func (h *UserHandler) SignUpHandler(c *gin.Context) {
	var info dto.SignUpDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	if err := validate.Struct(&info); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error.InvalidInput})
		return
	}

	err := h.s.SignUp(&info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

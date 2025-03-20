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
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	if err := validate.Struct(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput.Error())
		return
	}

	err := h.s.VerifyEmail(input.Email, input.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, error.VerifyError.Error())
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

package handler

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, error.BadRequest)
		return
	}

	if err := validate.Struct(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}

	err := h.s.VerifyEmail(input.Email, input.Code)
	if err != nil {
		switch {
		case errors.Is(err, error.VerificationCodeNotFound.Message):
			c.JSON(http.StatusBadRequest, error.VerificationCodeNotFound)
		case errors.Is(err, error.InvalidVerificationCode.Message):
			c.JSON(http.StatusBadRequest, error.InvalidVerificationCode)
		case errors.Is(err, error.ExpiredVerificationCode.Message):
			c.JSON(http.StatusBadRequest, error.ExpiredVerificationCode)
		default:
			c.JSON(http.StatusBadRequest, error.VerifyError)
		}
		return
	}
	c.JSON(http.StatusOK, server.Success)
}

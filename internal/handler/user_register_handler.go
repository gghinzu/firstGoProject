package handler

import (
	"errors"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"firstGoProject/internal/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *UserHandler) RegisterHandler(c *gin.Context) {
	var info dto.RegisterDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest)
		return
	}

	if !validation.ValidateEmail(info.Email) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidatePassword(info.Password) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateName(info.Name) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateSurname(info.Surname) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateAge(info.Age) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateGender(info.Gender) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateEducation(info.Education) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}

	err := h.s.Register(&info)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || err.Error() == "email is taken" {
			c.JSON(http.StatusConflict, error.AlreadyExists)
			return
		}
		c.JSON(http.StatusInternalServerError, error.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

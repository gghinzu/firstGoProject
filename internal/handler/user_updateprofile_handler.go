package handler

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"firstGoProject/internal/server"
	"firstGoProject/internal/validation"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"reflect"
)

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var updateData dto.UpdateProfileDTO

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest)
		return
	}

	if reflect.DeepEqual(updateData, dto.UpdateProfileDTO{}) {
		c.JSON(http.StatusBadRequest, error.EmptyRequestBody)
		return
	}

	if !validation.ValidateEmail(updateData.Email) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidatePassword(updateData.Password) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateName(updateData.Name) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateSurname(updateData.Surname) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateAge(updateData.Age) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateGender(updateData.Gender) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateStatus(updateData.Status) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateEducation(updateData.Education) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}
	if !validation.ValidateRole(updateData.Role.Name) {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput)
		return
	}

	claims := c.MustGet("claims").(helper.UserCustomClaims)

	roleInt := claims.Role

	if roleInt != string(enum.Admin) && updateData.Role.Name == enum.Admin {
		c.JSON(http.StatusUnauthorized, error.Unauthorized)
		return
	}

	_, err := h.s.UpdateProfile(claims.ID, &updateData)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, error.UpdateError)
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

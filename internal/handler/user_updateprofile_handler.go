package handler

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"firstGoProject/internal/server"
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
		c.JSON(http.StatusBadRequest, error.BadRequest)
		return
	}

	claims := c.MustGet("claims").(helper.UserCustomClaims)

	if claims.Role != string(enum.Admin) && updateData.Role.Name == enum.Admin {
		c.JSON(http.StatusUnprocessableEntity, error.Unauthorized)
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

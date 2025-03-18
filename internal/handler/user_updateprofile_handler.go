package handler

import (
	"errors"
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
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	if reflect.DeepEqual(updateData, dto.UpdateProfileDTO{}) {
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	claims := c.MustGet("claims").(helper.UserCustomClaims)

	_, err := h.s.UpdateProfile(claims.ID, &updateData)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, error.UpdateError.Error())
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

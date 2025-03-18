package handler

import (
	"errors"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *UserHandler) UpdateUserStatusByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, error.EmptyRequestBody.Error())
		return
	}

	var status *dto.StatusDTO
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusInternalServerError, error.JsonParseError.Error())
		return
	}

	var userStat enum.UserStatus
	userStatus := status.Status

	switch userStatus {
	case 0:
		userStat = enum.Active
	case 1:
		userStat = enum.Passive
	case 2:
		userStat = enum.Deleted
	default:
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	err := h.s.UpdateUserStatusByID(id, userStat)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound.Error())
			return
		}
		c.JSON(http.StatusNotFound, error.UpdateError.Error())
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

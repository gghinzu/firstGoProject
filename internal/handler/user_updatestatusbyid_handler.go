package handler

import (
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) UpdateUserStatusByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.EmptyBody})
		return
	}

	var status *dto.StatusDTO
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	err := h.s.UpdateUserStatusByID(id, userStat)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.NotFound})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

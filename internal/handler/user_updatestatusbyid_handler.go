package handler

import (
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UpdateUserStatusByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var status *dto.UserStatusDTO
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(500, gin.H{"error": "json cannot be bound"})
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
		c.JSON(400, gin.H{"error": "invalid status"})
	}

	err := h.s.UpdateUserStatusByID(id, userStat)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user status is successfully updated"})
}

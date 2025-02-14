package handler

import (
	"firstGoProject/internal/domain/enum"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *UserHandler) UpdateUserStatusByIDHandler(c *gin.Context) {
	id := c.Param("id")

	status := c.Param("status")
	var userStat enum.UserStatus
	userStatus, errStat := strconv.Atoi(status)
	if errStat != nil {
		c.JSON(400, gin.H{"error": "invalid status"})
	}

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
	}

	c.JSON(200, userStat)
}

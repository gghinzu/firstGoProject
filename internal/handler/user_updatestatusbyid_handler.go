package handler

import (
	"firstGoProject/internal/domain/enum"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *UserHandler) UpdateUserStatusByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	userStatus := c.Param("status")
	var userStat enum.UserStatus
	switch userStatus {
	case "passivate":
		userStat = enum.Passive
	case "activate":
		userStat = enum.Active
	default:
		c.JSON(400, gin.H{"error": "invalid status"})
	}

	err = h.s.UpdateUserStatusByID(id, userStat)

	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
}

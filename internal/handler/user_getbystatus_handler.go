package handler

import (
	"firstGoProject/internal/domain/enum"
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *UserHandler) GetUsersByStatusHandler(c *gin.Context) {
	statusStr := strings.TrimPrefix(c.Param("status"), "/")
	var sts enum.UserStatus

	switch statusStr {
	case "active":
		sts = enum.Active
	case "passive":
		sts = enum.Passive
	case "deleted":
		sts = enum.Deleted
	default:
		c.JSON(400, gin.H{"error": "invalid status"})
	}

	users, err := h.s.GetUsersByStatus(sts)
	if err != nil {
		c.JSON(404, gin.H{"error": "no user found"})
	}

	c.JSON(200, users)
}

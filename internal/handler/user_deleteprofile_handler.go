package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *UserHandler) DeleteProfile(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	err := h.s.DeleteProfile(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"message": "profile is deleted successfully"})
}

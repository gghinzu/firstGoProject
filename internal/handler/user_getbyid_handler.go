package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// GetUserByIDHandler uses the s layer's methods using their instances
func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	idStr := strings.TrimPrefix(c.Param("id"), "/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
		return
	}

	user, _ := h.s.GetUserByID(id)

	c.JSON(200, user)
}

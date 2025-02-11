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
		c.JSON(404, gin.H{"error": "invalid id"})
		return
	}

	user, errUser := h.s.GetUserByID(id)
	if user == nil {
		c.JSON(404, gin.H{"error": "no user found"})
		return
	}
	if errUser != nil {
		c.JSON(500, gin.H{"message": errUser.Error()})
		return
	}

	c.JSON(200, user)
}

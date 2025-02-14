package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// GetUserByIDHandler uses the s layer's methods using their instances
func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	user, errUser := h.s.GetUserByID(id)
	if user == nil {
		c.JSON(404, gin.H{"error": "no user found"})
		return
	}
	if errUser != nil {
		c.JSON(500, gin.H{"error": errUser.Error()})
		return
	}

	c.JSON(200, user)
}

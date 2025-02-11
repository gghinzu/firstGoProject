package handler

import (
	"github.com/gin-gonic/gin"
)

// GetUsersHandler uses the s layer's methods using their instances
func (h *UserHandler) GetUsersHandler(c *gin.Context) {
	users, err := h.s.GetAllUsers()
	if users == nil {
		c.JSON(404, gin.H{"error": "no users found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}
	c.JSON(200, users)
}

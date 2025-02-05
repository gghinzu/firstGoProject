package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	idStr := strings.TrimPrefix(c.Param("id"), "/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.s.DeleteUserByID(id)

	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

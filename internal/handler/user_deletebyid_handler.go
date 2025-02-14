package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	err := h.s.DeleteUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *UserHandler) DeleteProfile(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	err := h.s.DeleteProfile(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile is deleted successfully"})
}

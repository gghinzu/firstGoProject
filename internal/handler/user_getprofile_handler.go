package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	user, err := h.s.GetProfile(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

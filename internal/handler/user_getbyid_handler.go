package handler

import (
	"firstGoProject/internal/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	user, err := h.s.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.NotFound})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, user)
}

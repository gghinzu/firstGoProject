package handler

import (
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	err := h.s.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

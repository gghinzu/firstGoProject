package handler

import (
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) DeleteProfile(c *gin.Context) {
	claims := c.MustGet("claims").(helper.UserCustomClaims)

	err := h.s.DeleteProfile(claims.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

package handler

import (
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(helper.UserCustomClaims)

	user, err := h.s.GetProfile(claims.ID)
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

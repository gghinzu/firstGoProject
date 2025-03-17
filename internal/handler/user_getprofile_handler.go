package handler

import (
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(helper.UserCustomClaims)

	user, err := h.s.GetProfile(claims.ID)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": error.NotFound})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, user)
}

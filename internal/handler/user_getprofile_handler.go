package handler

import (
	"errors"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *UserHandler) GetProfile(c *gin.Context) {
	claims := c.MustGet("claims").(helper.UserCustomClaims)

	user, err := h.s.GetProfile(claims.ID)
	if user == nil {
		c.JSON(http.StatusNotFound, error.NotFound)
		return
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound)
			return
		}
		c.JSON(http.StatusInternalServerError, error.InternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

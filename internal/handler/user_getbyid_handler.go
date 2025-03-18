package handler

import (
	"errors"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	user, err := h.s.GetUserByID(id)
	if user == nil {
		c.JSON(http.StatusNotFound, server.NotFound)
		return
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, error.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

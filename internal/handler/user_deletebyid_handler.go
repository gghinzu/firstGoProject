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

func (h *UserHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	err := h.s.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, error.DeleteError)
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

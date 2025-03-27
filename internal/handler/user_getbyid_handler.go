package handler

import (
	"errors"
	"firstGoProject/internal/error"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	id := strings.TrimPrefix(c.Param("id"), "/")

	user, err := h.s.GetUserByID(id)
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

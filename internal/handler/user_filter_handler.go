package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) FilterHandler(c *gin.Context) {
	var info dto.FilterDTO
	err := c.ShouldBindQuery(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	users, err := h.s.FilterUser(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	if users == nil || len(*users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": error.NotFound})
		return
	}

	c.JSON(http.StatusOK, users)
}

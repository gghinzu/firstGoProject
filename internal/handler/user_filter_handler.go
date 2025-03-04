package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) FilterHandler(c *gin.Context) {
	var info dto.FilterDTO
	err := c.ShouldBindQuery(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := h.s.FilterUser(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if users == nil || len(*users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no user found"})
		return
	}

	c.JSON(http.StatusOK, users)
}

package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) FilterHandler(c *gin.Context) {
	var info dto.FilterDTO
	err := c.ShouldBindQuery(&info)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	users, errS := h.s.FilterUser(info)
	if errS != nil {
		c.JSON(500, gin.H{"message": errS.Error()})
		return
	}

	if users == nil || len(*users) == 0 {
		c.JSON(200, gin.H{"message": "no user found"})
		return
	}

	c.JSON(200, users)
}

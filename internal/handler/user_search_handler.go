package handler

import (
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
)

// SearchHandler gets url parameters and sends them to DTO
func (h *UserHandler) SearchHandler(c *gin.Context) {
	var info dto.SearchUserDTO
	err := c.ShouldBindQuery(&info)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	users, errS := h.s.SearchUser(info)
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

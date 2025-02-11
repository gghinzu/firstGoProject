package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (h *UserHandler) SearchHandler(c *gin.Context) {
	searchStr := strings.TrimPrefix(c.Param("search"), "/")

	users, err := h.s.SearchUser(searchStr)

	if users == nil {
		c.JSON(404, gin.H{"message": "no user found"})
	}
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
	}

	c.JSON(200, users)
}

package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (h *UserHandler) SoftDeleteUserByIDHandler(c *gin.Context) {
	idStr := strings.TrimPrefix(c.Param("id"), "/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "id should be numeric"})
		return
	}

	err = h.s.SoftDeleteUserByID(id)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
}

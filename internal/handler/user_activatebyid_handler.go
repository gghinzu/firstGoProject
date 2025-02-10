package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *UserHandler) ActivateUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser *entity.ActivatePassivateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	err = h.s.ActivateUserByID(id, updatedUser)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// OK
	c.JSON(200, gin.H{"message": "user is activated successfully"})
}

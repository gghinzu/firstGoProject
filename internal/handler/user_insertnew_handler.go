package handler

import (
	"firstGoProject/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (h *UserHandler) InsertNewUserHandler(c *gin.Context) {

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "Request body is empty"})
		return
	}

	var userModel *entity.UserDTO

	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.service.InsertNewUser(userModel)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User is added successfully"})

}

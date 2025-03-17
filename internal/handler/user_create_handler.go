package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.BadRequest})
		return
	}

	var userModel *dto.CreateDTO

	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.CreateUser(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

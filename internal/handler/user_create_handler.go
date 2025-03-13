package handler

import (
	"firstGoProject/internal/domain/error"
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
		return
	}

	var userModel *dto.CreateDTO

	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.CreateUser(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": error.Success})
}

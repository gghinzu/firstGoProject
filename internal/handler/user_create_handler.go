package handler

import (
	"errors"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, error.EmptyRequestBody)
		return
	}

	var userModel *dto.CreateDTO

	if err := c.ShouldBind(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, error.JsonParseError)
		return
	}

	err := h.s.CreateUser(userModel)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusForbidden, error.AlreadyExists)
			return
		}
		c.JSON(http.StatusInternalServerError, error.CreateError)
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

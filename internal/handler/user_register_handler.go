package handler

import (
	"errors"
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	validate = validator.New()
}

func (h *UserHandler) RegisterHandler(c *gin.Context) {
	var info dto.RegisterDTO

	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, error.BadRequest.Error())
		return
	}

	if err := validate.Struct(&info); err != nil {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput.Error())
		return
	}

	err := h.s.Register(&info)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusNotFound, error.AlreadyExists.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, error.InternalServerError.Error())
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

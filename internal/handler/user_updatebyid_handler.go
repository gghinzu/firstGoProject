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

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, error.EmptyRequestBody.Error())
		return
	}

	var updatedUser dto.UpdateDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, error.JsonParseError.Error())
		return
	}

	err := validate.Struct(&updatedUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, error.InvalidInput.Error())
		return
	}

	err = h.s.UpdateUserByID(id, &updatedUser)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, error.NotFound.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, error.UpdateError.Error())
		return
	}

	c.JSON(http.StatusOK, server.Success)
}

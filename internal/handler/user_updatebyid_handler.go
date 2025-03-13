package handler

import (
	"errors"
	"firstGoProject/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func init() {
	validate = validator.New()
}

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser dto.UpdateDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
		return
	}

	err := validate.Struct(&updatedUser)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": errs.Error()})
		return
	}

	err = h.s.UpdateUserByID(id, &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user is updated successfully"})
}

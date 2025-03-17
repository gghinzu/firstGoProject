package handler

import (
	"firstGoProject/internal/dto"
	"firstGoProject/internal/error"
	"firstGoProject/internal/server"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": error.EmptyBody})
		return
	}

	var updatedUser dto.UpdateDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json cannot be bound"})
		return
	}

	err := validate.Struct(&updatedUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": error.InvalidInput})
		return
	}

	err = h.s.UpdateUserByID(id, &updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"server": error.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": server.Success})
}

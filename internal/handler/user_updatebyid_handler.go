package handler

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (h *UserHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	auth, exist := c.Get("user")
	if !exist || auth == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user is unauthorized"})
	}

	currentUser := auth.(*entity.User)
	uid := uuid.MustParse(id)

	if currentUser.Role.Name != enum.Admin && currentUser.ID != uid {
		c.JSON(400, gin.H{"error": "this user can update only their own profile"})
		return
	}

	if c.Request.Body == nil {
		c.JSON(400, gin.H{"error": "request body is empty"})
		return
	}

	var updatedUser *entity.UpdateUserDTO

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(500, gin.H{"error": "json cannot be bound"})
		return
	}

	err := h.s.UpdateUserByID(id, updatedUser)
	if err != nil {
		c.JSON(404, gin.H{"error": "user cannot be found & updated"})
		return
	}
}

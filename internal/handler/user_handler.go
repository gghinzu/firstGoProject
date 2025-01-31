package handler

import (
	"firstGoProject/internal/domain/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// UserHandler is a structure to provide a communication with other layers
// we use structures for this purpose because we don't have 'implements' keyword in Go, for interface applications
// it communicates with the service
type UserHandler struct {
	service service.UserServicePort
}

// NewUserHandler provides an instance for interface
func NewUserHandler(service service.UserServicePort) *UserHandler {
	return &UserHandler{service: service}
}

// GetUsersHandler uses the service layer's methods using their instances
func (h *UserHandler) GetUsersHandler(c *gin.Context) {
	users := h.service.GetAllUsers()
	c.JSON(200, users)
}

// GetUserByIDHandler uses the service layer's methods using their instances
func (h *UserHandler) GetUserByIDHandler(c *gin.Context) {
	idStr := strings.TrimPrefix(c.Param("id"), "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

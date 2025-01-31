package handler

import (
	"github.com/gin-gonic/gin"
)

// GetUsersHandler uses the service layer's methods using their instances
func (h *UserHandler) GetUsersHandler(c *gin.Context) {
	users := h.service.GetAllUsers()
	c.JSON(200, users)
}

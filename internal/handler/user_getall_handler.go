package handler

import (
	"github.com/gin-gonic/gin"
)

// GetUsersHandler uses the s layer's methods using their instances
func (h *UserHandler) GetUsersHandler(c *gin.Context) {
	users := h.s.GetAllUsers()
	c.JSON(200, users)
}

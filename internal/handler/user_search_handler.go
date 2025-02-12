package handler

import "github.com/gin-gonic/gin"

func (h *UserHandler) SearchHandler(c *gin.Context) {
	name := c.Query("name")
	status := c.Query("status")
	gender := c.Query("gender")

	users, err := h.s.SearchUser(name, status, gender)

	if users == nil {
		c.JSON(404, gin.H{"message": "no user found"})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, users)
}

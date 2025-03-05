package middleware

import (
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(helper.UserCustomClaims)
		role := claims.Role

		if role != enum.Admin {
			c.JSON(http.StatusForbidden, gin.H{"error": "access is forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

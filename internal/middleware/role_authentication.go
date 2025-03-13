package middleware

import (
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(helper.UserCustomClaims)
		role := claims.Role

		if role != enum.Admin {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "access is forbidden"})
			return
		}
		c.Next()
	}
}

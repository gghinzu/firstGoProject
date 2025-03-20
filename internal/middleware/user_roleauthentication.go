package middleware

import (
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoleAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.MustGet("claims").(helper.UserCustomClaims)
		role := claims.Role

		if role != string(enum.Admin) {
			c.AbortWithStatusJSON(http.StatusForbidden, error.Unauthorized.Error())
			return
		}
		c.Next()
	}
}

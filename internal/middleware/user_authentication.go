package middleware

import (
	"firstGoProject/internal/error"
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, error.Unauthorized)
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := helper.VerifyToken(tokenString, "access")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, error.Unauthorized)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

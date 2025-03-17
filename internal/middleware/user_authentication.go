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
			c.JSON(http.StatusUnauthorized, gin.H{"error": error.NotAuthorized})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := helper.VerifyToken(tokenString, "access")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": error.Forbidden})
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

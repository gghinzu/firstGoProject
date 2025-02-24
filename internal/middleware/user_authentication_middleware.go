package middleware

import (
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		tokenErr := helper.VerifyToken(tokenString)
		if tokenErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": tokenErr.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

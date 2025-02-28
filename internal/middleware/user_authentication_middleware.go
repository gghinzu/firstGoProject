package middleware

import (
	"firstGoProject/internal/helper"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, tokenErr := helper.VerifyToken(tokenString)
		if tokenErr != nil {
			c.JSON(401, gin.H{"error": tokenErr.Error()})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

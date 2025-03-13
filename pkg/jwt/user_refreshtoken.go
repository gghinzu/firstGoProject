package jwt

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func init() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	secretKey = configuration.JWTRefreshSecret
}

func CreateRefreshToken(user *entity.User, role enum.UserRole) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    string(role),
		"type":    "refresh",
		"exp":     time.Now().Add(time.Hour * 24 * 180).Unix(),
	})

	secretKey := []byte(GetSecretKey())

	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

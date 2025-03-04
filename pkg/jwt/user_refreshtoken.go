package jwt

import (
	"firstGoProject/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
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

func CreateRefreshToken(userID uuid.UUID) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID.String(),
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

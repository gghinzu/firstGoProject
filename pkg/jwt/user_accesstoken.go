package jwt

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var secretKey string

func init() {
	configuration, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	secretKey = configuration.JWTAccessSecret
}

func CreateAccessToken(user *entity.User, role enum.UserRole) (*dto.TokenUserDTO, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID.String(),
		"role": string(role),
		"type": "access",
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	secretKey := []byte(GetSecretKey())

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &dto.TokenUserDTO{
		UserID:  user.ID,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Token:   accessTokenString,
	}, nil
}

func GetSecretKey() string {
	return secretKey
}

package jwt

import (
	"firstGoProject/internal/domain/entity"
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

func CreateAccessToken(userGiven *entity.User) (*dto.TokenUserDTO, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userGiven.ID.String(),
		"type": "access",
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	secretKey := []byte(GetSecretKey())

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &dto.TokenUserDTO{
		UserID:  userGiven.ID,
		Name:    userGiven.Name,
		Surname: userGiven.Surname,
		Email:   userGiven.Email,
		Token:   accessTokenString,
	}, nil
}

func GetSecretKey() string {
	return secretKey
}

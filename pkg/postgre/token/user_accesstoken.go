package token

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

var secretKey string

func init() {
	configuration, errC := config.LoadConfig()
	if errC != nil {
		log.Fatal("cannot load config:", errC)
	}
	secretKey = configuration.JWTSecret
}

func CreateAccessToken(user *entity.User) (*entity.TokenUserDTO, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})

	secretKey := []byte(GetSecretKey())

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &entity.TokenUserDTO{
		Email: &user.Email,
		Token: &accessTokenString,
	}, nil
}

func GetSecretKey() string {
	return secretKey
}

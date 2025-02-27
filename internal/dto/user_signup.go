package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type SignUpDTO struct {
	Email     string          `validate:"required" json:"email" gorm:"unique;not null"`
	Password  string          `validate:"required" json:"password" gorm:"not null"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       time.Time       `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

func (SignUpDTO) SignUpConvertToUser(dto *SignUpDTO) *entity.User {
	age := CalculateAge(dto.Age)

	userSign := &entity.User{
		Email:     dto.Email,
		Password:  dto.Password,
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}
	return userSign
}

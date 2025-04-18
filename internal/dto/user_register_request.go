package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type RegisterDTO struct {
	Email     string             `json:"email" gorm:"unique;not null"`
	Password  string             `json:"password" gorm:"not null"`
	Name      string             `json:"name"`
	Surname   string             `json:"surname"`
	Age       time.Time          `json:"age"`
	Gender    enum.UserGender    `json:"gender"`
	Education enum.UserEducation `json:"education"`
}

func (RegisterDTO) RegisterConvertToUser(dto *RegisterDTO) *entity.User {
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

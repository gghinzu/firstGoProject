package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

// TODO: regex
type RegisterDTO struct {
	Email     string             `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password  string             `json:"password" gorm:"not null" validate:"required,min=6,max=255"`
	Name      string             `json:"name" validate:"required,excludesall=0123456789,min=2,max=255"`
	Surname   string             `json:"surname" validate:"required,excludesall=0123456789,min=2,max=255"`
	Age       time.Time          `json:"age" validate:"required"`
	Gender    enum.UserGender    `json:"gender" validate:"required,oneof=male female 'not specified'"`
	Education enum.UserEducation `json:"education" validate:"required,oneof=None 'Primary School' 'Middle School' 'High School' 'Bachelor''s Degree' 'Master''s Degree' Doctorate"`
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

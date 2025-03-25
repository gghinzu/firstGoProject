package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type UpdateDTO struct {
	Name      string             `json:"name" validate:"required,excludesall=0123456789,min=2,max=255"`
	Surname   string             `json:"surname" validate:"required,excludesall=0123456789,min=2,max=255"`
	Age       time.Time          `json:"age" validate:"required"`
	Gender    enum.UserGender    `json:"gender" validate:"required,oneof=male female 'not specified'"`
	Education enum.UserEducation `json:"education" validate:"required,oneof=None 'Primary School' 'Middle School' 'High School' 'Bachelor''s Degree' 'Master''s Degree' Doctorate"`
}

func (UpdateDTO) UpdateConvertToUser(dto *UpdateDTO) *entity.User {
	age := CalculateAge(dto.Age)

	userUpdate := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       age,
		Gender:    dto.Gender,
		Education: dto.Education,
	}

	return userUpdate
}

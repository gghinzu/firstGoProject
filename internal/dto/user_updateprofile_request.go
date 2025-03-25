package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type UpdateProfileDTO struct {
	Email     string             `json:"email" validate:"required,email"`
	Password  string             `json:"password" validate:"required,min=6,max=255"`
	Name      string             `json:"name" validate:"required,excludesall=0123456789,min=2,max=255"`
	Surname   string             `json:"surname" validate:"required,excludesall=0123456789,min=2,max=255"`
	Age       time.Time          `json:"age" validate:"required"`
	Gender    enum.UserGender    `json:"gender" validate:"required,oneof=male female 'not specified'"`
	Status    enum.UserStatus    `json:"status" validate:"required,oneof=1 2 3"`
	Education enum.UserEducation `json:"education" validate:"required,oneof=None 'Primary School' 'Middle School' 'High School' 'Bachelor''s Degree' 'Master''s Degree' Doctorate"`
	Role      entity.UserRole    `json:"role" validate:"required,oneof=user admin"`
}

func (UpdateProfileDTO) UpdateProfileConvertToUser(dto *UpdateProfileDTO) *entity.User {
	age := CalculateAge(dto.Age)

	userUpdate := &entity.User{
		Email:     dto.Email,
		Password:  dto.Password,
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       age,
		Gender:    dto.Gender,
		Status:    dto.Status,
		Education: dto.Education,
		Role:      dto.Role,
	}

	return userUpdate
}

package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type UpdateProfileDTO struct {
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       time.Time       `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Role      entity.UserRole `json:"role"`
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
		Education: dto.Education,
		Role:      dto.Role,
	}
	return userUpdate
}

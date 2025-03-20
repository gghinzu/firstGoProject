package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type CreateDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       time.Time       `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Role      string          `json:"role"`
}

func (CreateDTO) CreateConvertToUser(dto *CreateDTO) *entity.User {
	age := CalculateAge(dto.Age)

	userCreate := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       age,
		Gender:    dto.Gender,
		Education: dto.Education,
		RoleID:    dto.Role,
	}

	return userCreate
}

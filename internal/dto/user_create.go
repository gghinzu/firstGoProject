package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

type CreateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       time.Time       `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Role      string          `json:"role"`
}

func (CreateUserDTO) CreateConvertToUser(dto *CreateUserDTO) *entity.User {
	age := CalculateAge(dto.Age)
	userCreate := &entity.User{
		Name:      dto.Name,
		Surname:   dto.Surname,
		Age:       age,
		Gender:    dto.Gender,
		Education: dto.Education,
		RoleID:    uuid.UUID{},
	}

	return userCreate
}

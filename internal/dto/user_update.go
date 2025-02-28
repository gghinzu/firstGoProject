package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type UpdateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       time.Time       `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

func (UpdateUserDTO) UpdateConvertToUser(dto *UpdateUserDTO) *entity.User {
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

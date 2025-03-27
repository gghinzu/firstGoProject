package dto

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"time"
)

type UpdateDTO struct {
	Name      string             `json:"name"`
	Surname   string             `json:"surname"`
	Age       time.Time          `json:"age"`
	Gender    enum.UserGender    `json:"gender"`
	Education enum.UserEducation `json:"education"`
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

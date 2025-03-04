package dto

import (
	"firstGoProject/internal/domain/enum"
)

type FilterDTO struct {
	//order
	Name      *string          `json:"name" form:"name"`
	Surname   *string          `json:"surname" form:"surname"`
	Age       *int             `json:"age" form:"age"`
	Gender    *enum.UserGender `json:"gender" form:"gender"`
	Education *string          `json:"education" form:"education"`
	Status    *enum.UserStatus `json:"status" form:"status"`
	Page      int32            `json:"page" form:"page"`
	Limit     int32            `json:"limit" form:"limit"`
}

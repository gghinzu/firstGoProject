package entity

import "firstGoProject/internal/domain/enum"

// for user creation, DTO
type CreateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

// for updating user, DTO
type UpdateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
}

// should bind (form)
type SearchUserDTO struct {
	Name      *string          `json:"name" form:"name"`
	Surname   *string          `json:"surname" form:"surname"`
	Age       *int             `json:"age" form:"age"`
	Gender    *enum.UserGender `json:"gender" form:"gender"`
	Education *string          `json:"education" form:"education"`
	Status    *enum.UserStatus `json:"status" form:"status"`
}

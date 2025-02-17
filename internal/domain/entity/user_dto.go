package entity

import (
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

// for user creation, DTO
type CreateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	RoleId    uuid.UUID       `json:"role_id"`
}

// for updating user, DTO
type UpdateUserDTO struct {
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Role      UserRole        `json:"role"`
}

// should bind (form) usage
type SearchUserDTO struct {
	Name      *string          `json:"name" form:"name"`
	Surname   *string          `json:"surname" form:"surname"`
	Age       *int             `json:"age" form:"age"`
	Gender    *enum.UserGender `json:"gender" form:"gender"`
	Education *string          `json:"education" form:"education"`
	Status    *enum.UserStatus `json:"status" form:"status"`
	Page      *int             `json:"page" form:"page"`
	Limit     *int             `json:"limit" form:"limit"`
}

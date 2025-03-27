package dto

import (
	"firstGoProject/internal/domain/enum"
)

type UserResponseDTO struct {
	ID        string          `json:"id"`
	Email     string          `json:"email"`
	Name      string          `json:"name"`
	Surname   string          `json:"surname"`
	Age       int             `json:"age"`
	Gender    enum.UserGender `json:"gender"`
	Education string          `json:"education"`
	Status    enum.UserStatus `json:"status"`
	Role      string          `json:"role,omitempty"`
}

package dto

import "github.com/google/uuid"

type TokenUserDTO struct {
	UserID       uuid.UUID `json:"user_id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
}

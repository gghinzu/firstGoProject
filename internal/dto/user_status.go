package dto

import "firstGoProject/internal/domain/enum"

type UserStatusDTO struct {
	Status enum.UserStatus `json:"status"`
}

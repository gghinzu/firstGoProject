package dto

import "firstGoProject/internal/domain/enum"

type StatusDTO struct {
	Status enum.UserStatus `json:"status"`
}

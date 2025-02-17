package entity

import (
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

type UserRole struct {
	RoleId uuid.UUID     `json:"r_id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	Name   enum.UserRole `json:"role_name" gorm:"not null"`
}

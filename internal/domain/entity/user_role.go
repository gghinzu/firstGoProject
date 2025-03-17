package entity

import (
	"firstGoProject/internal/domain/enum"
)

type UserRole struct {
	RoleId string        `json:"r_id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4();not null"`
	Name   enum.UserRole `json:"r_name" gorm:"not null"`
}

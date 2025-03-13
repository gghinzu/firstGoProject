package entity

import (
	"firstGoProject/internal/domain/enum"
)

type User struct {
	ID               string          `json:"id" gorm:"type:uuid;primaryKey;unique;default:uuid_generate_v4();not null"`
	Email            string          `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password         string          `json:"password" gorm:"type:varchar(255);not null"`
	Name             string          `json:"name" gorm:"type:varchar(255);not null"`
	Surname          string          `json:"surname" gorm:"type:varchar(255);not null"`
	Age              int             `json:"age" gorm:"type:bigint;not null"`
	Gender           enum.UserGender `json:"gender" gorm:"type:varchar(12);not null"`
	Education        string          `json:"education" gorm:"type:varchar(255);not null"`
	Status           enum.UserStatus `json:"status" gorm:"not null"`
	RoleID           string          `json:"role_id" gorm:"type:uuid;not null"`
	Role             UserRole        `json:"role" gorm:"not null;foreignKey:RoleID;references:RoleId"`
	VerificationCode string          `json:"verification_code" gorm:"type:varchar(255);not null"`
}

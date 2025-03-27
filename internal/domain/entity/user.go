package entity

import (
	"firstGoProject/internal/domain/enum"
	"time"
)

type User struct {
	ID                     string             `json:"id" gorm:"primaryKey;unique;default:uuid_generate_v4();not null"`
	Email                  string             `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password               string             `json:"password" gorm:"type:varchar(255);not null"`
	Name                   string             `json:"name" gorm:"type:varchar(255);not null"`
	Surname                string             `json:"surname" gorm:"type:varchar(255);not null"`
	Age                    int                `json:"age" gorm:"type:bigint;not null"`
	Gender                 enum.UserGender    `json:"gender" gorm:"not null"`
	Education              enum.UserEducation `json:"education" gorm:"not null"`
	Status                 enum.UserStatus    `json:"status" gorm:"not null"`
	RoleID                 string             `json:"role_id" gorm:"not null"`
	Role                   UserRole           `json:"role" gorm:"foreignKey:RoleID;references:RoleId;not null"`
	VerificationCode       *string            `json:"verification_code" gorm:"type:varchar(255)"`
	VerificationCodeExpiry *time.Time         `json:"verification_code_expiry" gorm:"type:timestamp"`
}

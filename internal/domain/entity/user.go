package entity

import (
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;unique;default:uuid_generate_v4()"`
	Email     string          `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password  string          `json:"password" gorm:"size:100;not null"`
	Name      string          `json:"name" gorm:"type:varchar(255);not null"`
	Surname   string          `json:"surname" gorm:"type:varchar(255);not null"`
	Age       int             `json:"age" gorm:"default:-1"`
	Gender    enum.UserGender `json:"gender" gorm:"type:varchar(12)"`
	Education string          `json:"education" gorm:"type:varchar(255)"`
	Status    enum.UserStatus `json:"status" gorm:"not null;default:0"`
	RoleID    uuid.UUID       `json:"role_id" gorm:"type:uuid;not null"`
	Role      UserRole        `json:"role" gorm:"not null;foreignKey:RoleID;references:RoleId"`
	/*	AccessToken  string          `json:"access_token"`
		RefreshToken string          `json:"refresh_token"`*/
}

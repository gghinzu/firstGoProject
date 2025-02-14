package entity

import (
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;unique;default:uuid_generate_v4()"`
	Name      string          `json:"name" gorm:"not null"`
	Surname   string          `json:"surname" gorm:"not null"`
	Age       int             `json:"age" gorm:"not null;default:-1"`
	Gender    enum.UserGender `json:"gender" gorm:"not null"`
	Education string          `json:"education" gorm:"not null"`
	Status    enum.UserStatus `json:"status" gorm:"not null;default:0"`
}

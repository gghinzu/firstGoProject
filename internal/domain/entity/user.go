package entity

import "firstGoProject/internal/domain/enum"

type User struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name      string          `json:"name" gorm:"not null"`
	Surname   string          `json:"surname" gorm:"not null"`
	Age       int             `json:"age" gorm:"not null;default:-1"`
	Gender    enum.UserGender `json:"gender" gorm:"not null"`
	Education string          `json:"education" gorm:"not null"`
	Status    enum.UserStatus `json:"status" gorm:"not null;default:0"`
}

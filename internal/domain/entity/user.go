package entity

import "firstGoProject/internal/domain/enum"

type User struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name      string          `json:"name" gorm:"not null"`
	Surname   string          `json:"surname" gorm:"not null"`
	Age       int             `json:"age" gorm:"not null"`
	Gender    string          `json:"gender" gorm:"not null"`
	Education string          `json:"education" gorm:"not null"`
	Status    enum.UserStatus `json:"status" gorm:"not null;default:1"`
	Deleted   bool            `json:"deleted" gorm:"default:false"`
}

// for user creation, DTO
type CreateUserDTO struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}

// for updating user, DTO
type UpdateUserDTO struct {
	Name      string `json:"name" `
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}

package repository

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
)

// UserRepositoryPort will be used as a bridge with other dependencies
type UserRepositoryPort interface {
	GetAllUsers() *[]entity.User
	GetUserByID(id int) (*entity.User, error)
	DeleteUserByID(id int) error
	UpdateUserByID(id int, updatedUser *entity.User) error
	CreateUser(newUser *entity.User) error
	GetUsersByStatus(status enum.UserStatus) (*[]entity.User, error)
	UpdateUserStatusByID(id int, userStatus enum.UserStatus) error
	SearchUser(searchStr string) (*[]entity.User, error)
}

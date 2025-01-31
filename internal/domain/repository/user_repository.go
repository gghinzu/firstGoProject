package repository

import "firstGoProject/internal/domain/entity"

// UserRepositoryPort will be used as a bridge with other dependencies
type UserRepositoryPort interface {
	GetAllUsers() []entity.User
	GetUserByID(id int) (*entity.User, error)
	DeleteUserByID(ID int) error
	UpdateUserByID(ID int, newUser entity.User) error
}

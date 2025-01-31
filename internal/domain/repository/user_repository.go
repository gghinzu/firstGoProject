package repository

import "firstGoProject/internal/domain/entity"

// UserRepositoryPort will be used as a bridge with other dependencies
type UserRepositoryPort interface {
	FetchAllUsers() []entity.User
	FetchUserByID(id int) (*entity.User, error)
}

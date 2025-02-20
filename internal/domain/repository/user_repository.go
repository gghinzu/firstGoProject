package repository

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"github.com/google/uuid"
)

// UserRepositoryPort will be used as a bridge with other dependencies
type UserRepositoryPort interface {
	GetUserByID(id uuid.UUID) (*entity.User, error)
	DeleteUserByID(id uuid.UUID) error
	UpdateUserByID(id uuid.UUID, updatedUser *entity.User) error
	CreateUser(newUser *entity.User) error
	UpdateUserStatusByID(id uuid.UUID, userStatus enum.UserStatus) error
	SearchUser(info entity.SearchUserDTO) (*[]entity.User, error)
}

package repository

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"github.com/google/uuid"
)

type UserRepositoryPort interface {
	GetUserByID(id uuid.UUID) (*entity.User, error)
	DeleteUserByID(id uuid.UUID) error
	UpdateUserByID(id uuid.UUID, updatedUser *entity.User) error
	CreateUser(newUser *entity.User) error
	UpdateUserStatusByID(id uuid.UUID, userStatus enum.UserStatus) error
	FilterUser(info dto.FilterDTO) (*[]entity.User, error)
	SignUp(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserRoleByRoleName(roleName string) (*entity.UserRole, error)
	GetUserWithRole(user *entity.User) (enum.UserRole, error)
}

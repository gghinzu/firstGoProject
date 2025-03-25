package repository

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
)

type UserRepositoryPort interface {
	GetUserByID(id string) (*dto.UserResponseDTO, error)
	GetUserByIDRaw(id string) (*entity.User, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, updatedUser *entity.User) error
	CreateUser(newUser *entity.User) error
	UpdateUserStatusByID(id string, userStatus enum.UserStatus) error
	FilterUser(info dto.FilterDTO) (*[]dto.UserResponseDTO, error)
	Register(user *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	GetUserRoleByRoleName(roleName string) (*entity.UserRole, error)
	GetRoleByUserInfo(user *entity.User) (enum.UserRole, error)
}

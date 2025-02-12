package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/domain/repository"
	"firstGoProject/pkg/postgre"
)

// UserServicePort is an interface, acts as a port to communicate with other layers
type UserServicePort interface {
	GetAllUsers() (*[]entity.User, error)
	GetUserByID(id int) (*entity.User, error)
	DeleteUserByID(id int) error
	UpdateUserByID(id int, updatedUser *entity.UpdateUserDTO) error
	CreateUser(newUser *entity.CreateUserDTO) error
	GetUsersByStatus(status enum.UserStatus) (*[]entity.User, error)
	UpdateUserStatusByID(id int, userStatus enum.UserStatus) error
	SearchUser(name, status, gender string) (*[]entity.User, error)
	SoftDeleteUserByID(id int) error
}

// UserService serves as a receiver for implementing UserRepositoryPort interface
type UserService struct {
	repo repository.UserRepositoryPort
}

// NewUserService to create a new instance, instead of directly initializing UserRepository
// instead of returning *UserRepository, it returns the interface UserRepositoryPort
// so that service layer doesn't need to know the exact struct
func NewUserService(repo *postgre.UserRepository) *UserService {
	return &UserService{repo: repo}
}

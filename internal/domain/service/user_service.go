package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/repository"
	"firstGoProject/pkg/postgres"
)

// UserServicePort is an interface, acts as a port to communicate with other layers
type UserServicePort interface {
	GetAllUsers() []entity.User
	GetUserByID(id int) *entity.User
	DeleteUserByID(id int) error
	UpdateUserByID(id int, updatedUser *entity.UserDTO) error
	InsertNewUser(newUser *entity.UserDTO) error
}

// UserService serves as a receiver for implementing UserRepositoryPort interface
type UserService struct {
	repo repository.UserRepositoryPort
}

// NewUserService to create a new instance, instead of directly initializing UserRepository
// instead of returning *UserRepository, it returns the interface UserRepositoryPort
// so that service layer doesn't need to know the exact struct
func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}

package service

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/repository"
)

// UserServicePort is an interface, acts as a port to communicate with other layers
type UserServicePort interface {
	GetAllUsers() []entity.User
	GetUserByID(id int) (*entity.User, error)
}

// GetAllUsers gets all users using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) GetAllUsers() []entity.User {
	return s.repo.FetchAllUsers()
}

// GetUserByID gets specified user with the given id using an instance of UserService
// (implementation of the interface UserServicePort)
func (s *UserService) GetUserByID(id int) (*entity.User, error) {
	user, err := s.repo.FetchUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UserService serves as a receiver for implementing UserRepositoryPort interface
type UserService struct {
	repo repository.UserRepositoryPort
}

// NewUserService to create a new instance, instead of directly initializing UserRepository
// instead of returning *UserRepository, it returns the interface UserRepositoryPort
// so that service layer doesn't need to know the exact struct
func NewUserService(repo repository.UserRepositoryPort) UserServicePort {
	return &UserService{repo: repo}
}

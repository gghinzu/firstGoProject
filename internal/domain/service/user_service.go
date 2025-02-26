package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/domain/repository"
	"firstGoProject/pkg/postgre"
)

// UserServicePort is an interface, acts as a port to communicate with other layers
type UserServicePort interface {
	GetUserByID(id string) (*entity.User, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, updatedUser *entity.UpdateUserDTO) error
	CreateUser(newUser *entity.CreateUserDTO) error
	UpdateUserStatusByID(id string, userStatus enum.UserStatus) error
	SearchUser(info entity.SearchUserDTO) (*[]entity.User, error)
	SignUp(newUser *entity.SignUpDTO) error
	Login(info entity.LoginDTO) (*entity.TokenUserDTO, error)
	RefreshToken(email *string) (*entity.TokenUserDTO, error)
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

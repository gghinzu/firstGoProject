package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/domain/repository"
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/postgres"
)

type UserServicePort interface {
	GetUserByID(id string) (*entity.User, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, updatedUser *dto.UpdateUserDTO) error
	CreateUser(newUser *dto.CreateUserDTO) error
	UpdateUserStatusByID(id string, userStatus enum.UserStatus) error
	FilterUser(info dto.FilterDTO) (*[]entity.User, error)
	SignUp(newUser *dto.SignUpDTO) error
	Login(info dto.LoginDTO) (*dto.TokenUserDTO, error)
	RefreshToken() (*dto.TokenUserDTO, error)
	GetProfile(id string) (*entity.User, error)
	UpdateProfile(id string, userDTO *dto.UpdateProfileDTO) (*entity.User, error)
	DeleteProfile(id string) error
}

type UserService struct {
	repo repository.UserRepositoryPort
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}

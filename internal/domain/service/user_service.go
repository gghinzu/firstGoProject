package service

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/domain/repository"
	"firstGoProject/internal/dto"
	"firstGoProject/pkg/postgres"
)

type UserServicePort interface {
	GetUserByID(id string) (*dto.UserResponseDTO, error)
	DeleteUserByID(id string) error
	UpdateUserByID(id string, updatedUser *dto.UpdateDTO) error
	UpdateUserStatusByID(id string, userStatus enum.UserStatus) error
	FilterUser(info dto.FilterDTO) (*[]dto.UserResponseDTO, error)
	Register(newUser *dto.RegisterDTO) error
	Login(info dto.LoginDTO) (*dto.TokenDTO, error)
	RefreshToken(id string) (*dto.TokenDTO, error)
	GetProfile(id string) (*dto.UserResponseDTO, error)
	UpdateProfile(id string, userDTO *dto.UpdateProfileDTO) (*entity.User, error)
	GetUserWithRole(id string) (*entity.User, enum.UserRole, error)
	VerifyEmail(email, code string) error
}

type UserService struct {
	repo repository.UserRepositoryPort
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{repo: repo}
}

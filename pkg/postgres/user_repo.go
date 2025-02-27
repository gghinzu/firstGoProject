package postgres

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	user2 "firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*entity.User, error) {
	var getUser *entity.User
	err := r.db.Preload("Role").Where("id = ?", id).First(&getUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return getUser, nil
}

func (r *UserRepository) DeleteUserByID(id uuid.UUID) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserByID(id uuid.UUID, updatedUser *entity.User) error {
	err := r.db.Where("id = ?", id).Updates(&updatedUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) CreateUser(newUser *entity.User) error {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserStatusByID(id uuid.UUID, userStatus user2.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FilterUser(info dto.FilterDTO) (*[]entity.User, error) {
	var users *[]entity.User
	query := r.db.Preload("Role")

	if info.Name != nil {
		query = query.Where("name ILIKE ?", info.Name)
	}
	if info.Surname != nil {
		query = query.Where("surname ILIKE ?", info.Surname)
	}
	if info.Education != nil {
		query = query.Where("education ILIKE ?", info.Education)
	}
	if info.Gender != nil {
		query = query.Where("gender ILIKE ?", info.Gender)
	}
	if info.Age != nil {
		query = query.Where("age = ?", info.Age)
	}
	if info.Status != nil {
		query = query.Where("status = ?", info.Status)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) SignUp(info *entity.User) error {
	err := r.db.Create(&info).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserRoleByRoleName(roleName string) (*entity.UserRole, error) {
	var userRole *entity.UserRole
	err := r.db.Model(&entity.UserRole{}).Where("name = ?", roleName).First(&userRole).Error
	if err != nil {
		return nil, err
	}
	return userRole, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var newUser *entity.User
	err := r.db.Where("email ILIKE ?", email).First(&newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

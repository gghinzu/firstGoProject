package postgres

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"firstGoProject/internal/dto"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetUserByID(id string) (*entity.User, error) {
	var getUser *entity.User
	err := r.db.Preload("Role").Where("id = ?", id).First(&getUser).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return getUser, nil
}

func (r *UserRepository) DeleteUserByID(id string) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserByID(id string, updatedUser *entity.User) error {
	if updatedUser == nil {
		return errors.New("updatedUser cannot be nil")
	}

	result := r.db.Where("id = ?", id).Select("*").Updates(*updatedUser)
	if result.Error != nil {
		return fmt.Errorf("failed to update user with id %s: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("no user found with the given id")
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

func (r *UserRepository) UpdateUserStatusByID(id string, userStatus enum.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FilterUser(info dto.FilterDTO) (*[]entity.User, error) {
	var users *[]entity.User
	query := r.paginate(int(info.Limit), int(info.Page)).Preload("Role")

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
	if info.Order != nil {
		query = query.Order(*info.Order)
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

func (r *UserRepository) GetUserWithRole(user *entity.User) (enum.UserRole, error) {
	var userRole entity.UserRole
	err := r.db.Where("role_id = ?", user.RoleID).First(&userRole).Error
	if err != nil {
		return "", err
	}
	return userRole.Name, nil
}

func (r *UserRepository) paginate(limit int, offset int) *gorm.DB {
	var users *[]entity.User
	return r.db.Limit(limit).Offset((offset - 1) * limit).Find(&users)
}

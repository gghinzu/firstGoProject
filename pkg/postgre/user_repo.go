package postgre

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/enum"
	"gorm.io/gorm"
)

// UserRepository is used since Go does not have explicit 'implements' keyword, an empty struct is used to implicitly implement interfaces
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

var gError = errors.New("record not found")

// GetAllUsers for displaying all the users
func (r *UserRepository) GetAllUsers() (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Order("id").Find(&userList).Error
	if err != nil {
		return nil, gError
	}
	return userList, nil
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, gError
	}
	return user, nil
}

// DeleteUserByID to delete a specific user by the given id
func (r *UserRepository) DeleteUserByID(id int) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return gError
	}
	return nil
}

// UpdateUserByID to update a specific user's info by the given id
func (r *UserRepository) UpdateUserByID(id int, updatedUser *entity.User) error {
	err := r.db.Where("id = ?", id).Updates(&updatedUser).Error
	if err != nil {
		return gError
	}
	return nil
}

// CreateUser to create a new user
func (r *UserRepository) CreateUser(newUser *entity.User) error {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// GetUsersByStatus lists all users according to their status
func (r *UserRepository) GetUsersByStatus(status enum.UserStatus) (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Order("id").Where("status = ?", status).Find(&userList).Error
	if err != nil {
		return nil, gError
	}
	return userList, nil
}

// UpdateUserStatusByID updates users' statuses
func (r *UserRepository) UpdateUserStatusByID(id int, userStatus enum.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return errors.New("user cannot be updated")
	}
	return nil
}

// SearchUser searches for a specific string in users' info
func (r *UserRepository) SearchUser(name, status, gender string) (*[]entity.User, error) {
	var userList *[]entity.User
	query := r.db.Where("lower(name) LIKE lower(?) OR lower(surname) LIKE lower(?)", "%"+name+"%", "%"+name+"%")
	if status != "" {
		query = query.Where("lower(status) = lower(?)", "%"+status+"%")
	} else if gender != "" {
		query = query.Where("lower(gender) LIKE lower(?)", "%"+gender+"%")
	}
	err := query.Find(&userList).Error
	if err != nil {
		return nil, gError
	}
	return userList, nil
}

func (r *UserRepository) SoftDeleteUserByID(id int) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("deleted", true).Error
	if err != nil {
		return gError
	}
	return nil
}

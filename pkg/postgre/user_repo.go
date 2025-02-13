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

// GetAllUsers for displaying all the users
func (r *UserRepository) GetAllUsers() (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Order("id").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return user, nil
}

// DeleteUserByID to delete a specific user by the given id
func (r *UserRepository) DeleteUserByID(id int) error {
	err := r.db.Delete(&entity.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserByID to update a specific user's info by the given id
func (r *UserRepository) UpdateUserByID(id int, updatedUser *entity.User) error {
	err := r.db.Where("id = ?", id).Updates(&updatedUser).Error
	if err != nil {
		return err
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
		return nil, err
	}
	return userList, nil
}

// UpdateUserStatusByID updates users' statuses
func (r *UserRepository) UpdateUserStatusByID(id int, userStatus enum.UserStatus) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
	if err != nil {
		return err
	}
	return nil
}

// SearchUser searches for a specific value in users' info and lists users
func (r *UserRepository) SearchUser(info *entity.SearchUserDTO) (*[]entity.User, error) {
	var users *[]entity.User
	query := r.db

	if info.Name != "" {
		query = query.Where("name ILIKE ?", "%"+info.Name+"%")
	}
	if info.Surname != "" {
		query = query.Where("surname ILIKE ?", "%"+info.Surname+"%")
	}
	if info.Education != "" {
		query = query.Where("education ILIKE ?", "%"+info.Education+"%")
	}
	if info.Gender != "" {
		query = query.Where("gender ILIKE ?", info.Gender)
	}
	if info.Age > 0 {
		query = query.Where("age = ?", info.Age)
	}
	if info.Status != enum.NotInitialized {
		query = query.Where("status = ?", info.Status)
	}

	err := query.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// SoftDeleteUserByID disables user and their chances except their statuses
func (r *UserRepository) SoftDeleteUserByID(id int) error {
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", enum.Deleted).Error
	if err != nil {
		return err
	}
	return nil
}

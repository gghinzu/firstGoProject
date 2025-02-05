package postgre

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"gorm.io/gorm"
)

// UserRepository is used since Go does not have explicit 'implements' keyword, an empty struct is used to implicitly implement interfaces
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAllUsers for displaying all the users
func (r *UserRepository) GetAllUsers() *[]entity.User {
	var userList *[]entity.User
	err := r.DB.Order("id").Find(&userList).Error
	if err != nil {
		panic(err)
	}
	return userList
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	var user *entity.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		} else {
			return nil, err
		}
	}
	return user, nil
}

// DeleteUserByID to delete a specific user by the given id
func (r *UserRepository) DeleteUserByID(id int) error {
	err := r.DB.Where("id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		} else {
			return err
		}
	} else {
		err = r.DB.Delete(&entity.User{}, id).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) UpdateUserByID(id int, updatedUser *entity.User) error {
	err := r.DB.Where("id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		} else {
			return err
		}
	} else {
		err = r.DB.Updates(updatedUser).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *UserRepository) CreateUser(newUser *entity.User) error {
	err := r.DB.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

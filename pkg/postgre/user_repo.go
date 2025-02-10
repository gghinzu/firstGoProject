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
func (r *UserRepository) GetAllUsers() *[]entity.User {
	var userList *[]entity.User
	err := r.db.Order("id").Find(&userList).Error
	if err != nil {
		panic(err)
	}
	return userList
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	var user *entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
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
	err := r.db.Where("id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		} else {
			return err
		}
	} else {
		err = r.db.Delete(&entity.User{}, id).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateUserByID to update a specific user's info by the given id
func (r *UserRepository) UpdateUserByID(id int, updatedUser *entity.User) error {
	err := r.db.Where("id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		} else {
			return err
		}
	} else {
		err = r.db.Where("id = ?", id).Updates(&updatedUser).Error
		if err != nil {
			return err
		}
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

func (r *UserRepository) GetUsersByStatus(status enum.UserStatus) (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Order("id").Where("status = ?", status).Find(&userList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
	}
	return userList, nil
}

func (r *UserRepository) ActivateUserByID(id int, updatedUser *entity.User) error {
	err := r.UpdateUserByID(id, updatedUser)
	if err != nil {
		return err
	}
	// using 'Update' explicitly because assignment with zero might be misunderstood in go
	err = r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", updatedUser.Status).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) PassivateUserByID(id int, updatedUser *entity.User) error {
	err := r.UpdateUserByID(id, updatedUser)
	if err != nil {
		return err
	}
	err = r.db.Model(&entity.User{}).Where("id = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

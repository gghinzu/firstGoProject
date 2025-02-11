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

// GetUsersByStatus lists all users according to their status
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

// UpdateUserStatusByID updates users' statuses
func (r *UserRepository) UpdateUserStatusByID(id int, userStatus enum.UserStatus) error {
	user, _ := r.GetUserByID(id)
	if user.Status == enum.Active && userStatus == enum.Active {
		err := errors.New("user is already active")
		return err
	} else if user.Status == enum.Passive && userStatus == enum.Passive {
		err := errors.New("user is already passive")
		return err
	} else {
		// using 'Update' explicitly because assignment with zero might be misunderstood in go
		err := r.db.Model(&entity.User{}).Where("id = ?", id).Update("status", userStatus).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// SearchUser searches for a specific string in users' names
func (r *UserRepository) SearchUser(searchString string) (*[]entity.User, error) {
	var userList *[]entity.User
	err := r.db.Where("name LIKE ?", "%"+searchString+"%").Or("surname LIKE ?", "%"+searchString+"%").Find(&userList).Error
	//err := r.db.Order("id").Where("(name, surname) IN ?", [][]interface{}{{searchString, searchString}}).Find(&userList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic("user not found")
		}
	}
	return userList, nil
}

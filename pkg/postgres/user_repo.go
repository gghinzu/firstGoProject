package postgres

import (
	"errors"
	"firstGoProject/internal/domain/entity"
)

// Users list with mock data
var Users = []entity.User{
	{1, "John", "Doe", 28, "Male", "Bachelor's in Computer Science"},
	{2, "Jane", "Smith", 25, "Female", "Master's in Civil Engineering"},
	{3, "Mike", "Johnson", 32, "Male", "PhD in Physics"},
	{4, "Emily", "Davis", 27, "Female", "Bachelor's in Business Administration"},
	{5, "Robert", "Brown", 40, "Male", "Master's in Mechanical Engineering"},
}

// UserRepository is used since Go does not have explicit 'implements' keyword, an empty struct is used to implicitly implement interfaces
type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// GetAllUsers for displaying all the users
func (r *UserRepository) GetAllUsers() []entity.User {
	return Users
}

// GetUserByID to get a specific user's details
func (r *UserRepository) GetUserByID(ID int) (*entity.User, error) {
	for _, user := range Users {
		if user.ID == ID {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

// DeleteUserByID to delete a specific user by the given id
func (r *UserRepository) DeleteUserByID(ID int) error {
	for index, user := range Users {
		if user.ID == ID {
			Users = append(Users[:index], Users[index+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) UpdateUserByID(id int, updatedUser entity.User) error {
	for i, user := range Users {
		if user.ID == id {
			Users[i] = updatedUser
			return nil
		}
	}
	return errors.New("user not found")
}

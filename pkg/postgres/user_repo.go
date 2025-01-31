package postgres

import (
	"errors"
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/repository"
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

func NewUserRepository() repository.UserRepositoryPort {
	return &UserRepository{}
}

// FetchAllUsers for displaying all the users
func (r *UserRepository) FetchAllUsers() []entity.User {
	return Users
}

// FetchUserByID to get a specific user's details
func (r *UserRepository) FetchUserByID(ID int) (*entity.User, error) {
	for _, user := range Users {
		if user.ID == ID {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}

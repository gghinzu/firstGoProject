package person

import (
	"errors"
)

// person structure, ok for json usage
type Person struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Education string `json:"education"`
}

// user list with mock data
var people = []Person{
	{1, "John", "Doe", 28, "Male", "Bachelor's in Computer Science"},
	{2, "Jane", "Smith", 25, "Female", "Master's in Civil Engineering"},
	{3, "Mike", "Johnson", 32, "Male", "PhD in Physics"},
	{4, "Emily", "Davis", 27, "Female", "Bachelor's in Business Administration"},
	{5, "Robert", "Brown", 40, "Male", "Master's in Mechanical Engineering"},
}

// function for displaying all the users
func GetAllUsers() []Person {
	return people
}

// function to get a specific user's details
func GetUserDetail(ID int) (Person, error) {
	for _, person := range people {
		if person.ID == ID {
			return person, nil
		}
	}
	return Person{}, errors.New("User not found")
}

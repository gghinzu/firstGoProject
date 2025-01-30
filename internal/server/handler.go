package server

import (
	"firstGoProject/internal/person"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

//**code 200 means ok, code 404 means error**

// we will get all the users and display them as they are with this function
// using gin framework
func GetUsersHandler(c *gin.Context) {
	people := person.GetAllUsers()
	c.JSON(http.StatusOK, people)
}

// we will get just only the given user & their details with this function
func GetUsersByIDHandler(c *gin.Context) {
	//trimming leading '/'
	//prevents errors in type conversion ('1' instead of '/1')
	idFromUrl := strings.TrimPrefix(c.Param("id"), "/")
	//converting given id to integer
	id, err := strconv.Atoi(idFromUrl)
	//handling error
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid ID"})
	}
	//fetching user details
	person, err := person.GetUserDetail(id)
	//handling error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
	}
	//return user details as json
	c.JSON(200, person)
}

package main

import (
	"firstGoProject/internal/domain/service"
	"firstGoProject/internal/handler"
	"firstGoProject/pkg/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// using gin framework
	router := gin.Default()

	// dependency injection
	repo := postgres.NewUserRepository()
	serv := service.NewUserService(repo)
	hand := handler.NewUserHandler(serv)

	// manipulating the url with gin
	router.GET("/user-list", hand.GetUsersHandler)
	// we are using '/' notation in url convention instead of other special characters
	// here, '*' is for capturing values for 'id'
	// it will be shown like that -> http://localhost:3000/get-user-by-id/1
	router.GET("/get-user-by-id/*id", hand.GetUserByIDHandler)

	// validation
	fmt.Println("Server is running on http://localhost:3000")

	// the specific port that we want to work with & error handling
	err := router.Run(":3000")
	if err != nil {
		return
	}
}

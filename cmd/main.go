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
	//grouping
	router := gin.Default()

	// dependency injection
	repo := postgres.NewUserRepository()

	userService := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(userService)

	// manipulating the url with gin
	router.GET("", userHandler.GetUsersHandler)
	// we are using '/' notation in url convention instead of other special characters
	// here, '*' is for capturing values for 'id'
	// it will be shown like that -> http://localhost:3000/get-user-by-id/1
	router.GET("/:id", userHandler.GetUserByIDHandler)
	router.DELETE("/*id", userHandler.DeleteUserByIDHandler)

	// validation
	fmt.Println("Server is running on http://localhost:3000")

	// the specific port that we want to work with & error handling
	err := router.Run(":3000")
	if err != nil {
		return
	}
}

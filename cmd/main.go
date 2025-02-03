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

	// dependency injections
	userRepository := postgres.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// router grouping
	users := router.Group("/users")
	{
		users.GET("", userHandler.GetUsersHandler)
		users.GET("/:id", userHandler.GetUserByIDHandler)
		users.DELETE("/:id", userHandler.DeleteUserByIDHandler)
		users.PUT("/:id", userHandler.UpdateUserByIDHandler)
		users.POST("", userHandler.InsertNewUserHandler)
	}

	// validation
	fmt.Println("Server is running on http://localhost:3000")

	// the specific port that we want to work with & error handling
	err := router.Run(":3000")
	if err != nil {
		return
	}
}

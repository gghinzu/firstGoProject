package main

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/repository/seed"
	"firstGoProject/internal/domain/service"
	"firstGoProject/internal/handler"
	"firstGoProject/pkg/postgre"
	"github.com/gin-gonic/gin"
)

func main() {
	// using gin framework
	router := gin.Default()

	// there's no need for close connection, gorm does it automatically
	db := postgre.Connection()
	err := seed.Seed(db)
	if err != nil {
		return
	}

	// dependency injections
	userRepository := postgre.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// migration, code first, database generator according to entity
	err = db.Migrator().AutoMigrate(&entity.User{})
	if err != nil {
		return
	}

	// router grouping
	users := router.Group("/user")
	{
		users.GET("/:id", userHandler.GetUserByIDHandler)
		users.GET("", userHandler.SearchHandler)
		users.PUT("/:id", userHandler.UpdateUserByIDHandler)
		//dto may be used for updating the statuses
		users.PUT("/:id/:status", userHandler.UpdateUserStatusByIDHandler)
		users.POST("", userHandler.CreateUserHandler)
		users.DELETE("/:id", userHandler.DeleteUserByIDHandler)
	}

	// the specific port that we want to work with & error handling
	err = router.Run(":3000")
	if err != nil {
		return
	}
}

package main

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/service"
	"firstGoProject/internal/handler"
	"firstGoProject/internal/middleware"
	"firstGoProject/pkg/config"
	"firstGoProject/pkg/postgre"
	"firstGoProject/pkg/postgre/seed"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configure, errC := config.LoadConfig()
	if errC != nil {
		log.Fatal("cannot load config:", errC)
	}

	// using gin framework
	router := gin.New()

	// there's no need for close connection, gorm does it automatically
	db := postgre.Connection()

	// dependency injections
	userRepository := postgre.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// migration, code first, database generator according to entity
	err := db.Migrator().AutoMigrate(&entity.User{}, &entity.UserRole{})
	if err != nil {
		log.Fatalf("auto migration failed: %s", err.Error())
		return
	}

	err = seed.Seed(db, &configure)
	if err != nil {
		log.Fatalf("seeding failed: %v\n", err)
		return
	}

	// router grouping
	users := router.Group("/user")
	{
		authorized := users.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			authorized.GET("/:id", userHandler.GetUserByIDHandler)
			authorized.GET("/:id/:status", userHandler.UpdateUserStatusByIDHandler)
			authorized.PUT("/:id", userHandler.UpdateUserByIDHandler)
			authorized.POST("", userHandler.CreateUserHandler)
			authorized.DELETE("/:id", userHandler.DeleteUserByIDHandler)
			authorized.GET("", userHandler.SearchHandler)
		}
		//dto may be used for updating the statuses
		users.POST("/register", userHandler.SignUpHandler)
		users.POST("/login", userHandler.LoginHandler)
		users.POST("/refresh-token", userHandler.RefreshTokenHandler)
	}

	// the specific port that we want to work with & error handling
	err = router.Run(configure.ClientOrigin)
	if err != nil {
		return
	}
}

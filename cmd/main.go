package main

import (
	"firstGoProject/internal/domain/entity"
	"firstGoProject/internal/domain/service"
	"firstGoProject/internal/handler"
	"firstGoProject/internal/middleware"
	"firstGoProject/pkg/config"
	"firstGoProject/pkg/postgres"
	"firstGoProject/pkg/postgres/seed"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	configure, errC := config.LoadConfig()
	if errC != nil {
		log.Fatal("cannot load config:", errC)
	}

	router := gin.New()

	db := postgres.Connection()

	// dependency injections
	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	err := db.Migrator().AutoMigrate(&entity.User{}, &entity.UserRole{})
	if err != nil {
		log.Fatalf("auto migration failed: %s", err.Error())
		return
	}

	errRoleSeed := seed.RoleSeed(db)
	if errRoleSeed != nil {
		log.Fatalf("role seeding failed: %v\n", errRoleSeed)
	}

	errUserSeed := seed.AdminSeed(db)
	if errUserSeed != nil {
		log.Fatalf("user seeding failed: %v\n", errUserSeed)
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
			authorized.GET("", userHandler.FilterHandler)
		}
		users.POST("/register", userHandler.SignUpHandler)
		users.POST("/login", userHandler.LoginHandler)
		users.POST("/refresh-token", userHandler.RefreshTokenHandler)
	}

	err = router.Run(configure.ClientOrigin)
	if err != nil {
		return
	}
}

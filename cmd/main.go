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
	configure, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	router := gin.New()

	db := postgres.Connection()

	// dependency injections
	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	err = db.Migrator().AutoMigrate(&entity.User{}, &entity.UserRole{})
	if err != nil {
		log.Fatalf("auto migration failed: %s", err.Error())
		return
	}

	err = seed.RoleSeed(db)
	if err != nil {
		log.Fatalf("role seeding failed: %v\n", err)
	}

	err = seed.AdminSeed(db)
	if err != nil {
		log.Fatalf("user seeding failed: %v\n", err)
		return
	}

	// router grouping
	users := router.Group("/user")
	{
		authorized := users.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			profile := authorized.Group("/:id/profile")
			{
				profile.GET("", userHandler.GetProfile)
				profile.PUT("", userHandler.UpdateProfile)
				profile.DELETE("", userHandler.DeleteProfile)
			}
			adminAuth := authorized.Group("")
			adminAuth.Use(middleware.RoleMiddleware())
			{
				adminAuth.GET("/:id", userHandler.GetUserByIDHandler)
				adminAuth.POST("/:id/status", userHandler.UpdateUserStatusByIDHandler)
				adminAuth.PUT("/:id", userHandler.UpdateUserByIDHandler)
				adminAuth.POST("", userHandler.CreateUserHandler)
				adminAuth.DELETE("/:id", userHandler.DeleteUserByIDHandler)
				adminAuth.GET("", userHandler.FilterHandler)
			}
		}
		users.POST("/refresh-token", userHandler.RefreshTokenHandler)
		users.POST("/register", userHandler.SignUpHandler)
		users.POST("/login", userHandler.LoginHandler)
	}

	err = router.Run(configure.ClientOrigin)
	if err != nil {
		return
	}
}

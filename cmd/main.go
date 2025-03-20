package main

import (
	"firstGoProject/internal/domain/service"
	"firstGoProject/internal/handler"
	"firstGoProject/internal/middleware"
	"firstGoProject/pkg/config"
	"firstGoProject/pkg/postgres"
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
	postgres.Migrate(db)
	postgres.Seed(db)

	userRepository := postgres.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// router grouping
	users := router.Group("/user")
	{
		authorized := users.Group("")
		authorized.Use(middleware.AuthMiddleware())
		{
			profile := authorized.Group("/profile")
			{
				profile.GET("", userHandler.GetProfile)
				profile.PUT("", userHandler.UpdateProfile)
				profile.DELETE("", userHandler.DeleteProfile)
			}
			adminAuth := authorized.Group("")
			adminAuth.Use(middleware.RoleAuthentication())
			{
				adminAuth.GET("", userHandler.FilterHandler)
				adminAuth.GET("/:id", userHandler.GetUserByIDHandler)
				adminAuth.PUT("/:id", userHandler.UpdateUserByIDHandler)
				adminAuth.POST("", userHandler.CreateUserHandler)
				adminAuth.POST("/:id/status", userHandler.UpdateUserStatusByIDHandler)
				adminAuth.DELETE("/:id", userHandler.DeleteUserByIDHandler)
			}
		}
		users.POST("/refresh-token", userHandler.RefreshTokenHandler)
		users.POST("/register", userHandler.RegisterHandler)
		users.POST("/login", userHandler.LoginHandler)
		users.POST("/verify", userHandler.VerifyEmailHandler)
	}

	err = router.Run(configure.ClientOrigin)
	if err != nil {
		return
	}
}

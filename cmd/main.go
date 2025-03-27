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

	SetupRouter(router, userHandler)

	err = router.Run(configure.ClientOrigin)
	if err != nil {
		return
	}
}

func SetupRouter(router *gin.Engine, userHandler *handler.UserHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/refresh-token", userHandler.RefreshTokenHandler)
		auth.POST("/login", userHandler.LoginHandler)
	}

	profile := router.Group("/profile")
	profile.Use(middleware.AuthMiddleware())
	{
		profile.GET("", userHandler.GetProfile)
		profile.PUT("", userHandler.UpdateProfile)
	}

	user := router.Group("/user")
	{
		user.POST("/register", userHandler.RegisterHandler)
		user.POST("/verify", userHandler.VerifyEmailHandler)
	}

	user.Use(middleware.AuthMiddleware(), middleware.RoleAuthorization())
	{
		user.GET("", userHandler.FilterHandler)
		user.GET("/:id", userHandler.GetUserByIDHandler)
		user.PUT("/:id", userHandler.UpdateUserByIDHandler)
		user.POST("/:id/status", userHandler.UpdateUserStatusByIDHandler)
		user.DELETE("/:id", userHandler.DeleteUserByIDHandler)
	}
}

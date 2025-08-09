package main

import (
	"github.com/ardianilyas/go-auth/config"
	"github.com/ardianilyas/go-auth/internal/handlers"
	"github.com/ardianilyas/go-auth/internal/middlewares"
	"github.com/ardianilyas/go-auth/internal/models"
	"github.com/ardianilyas/go-auth/internal/repositories"
	"github.com/ardianilyas/go-auth/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open(config.DB_DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	userRepo := &repositories.UserRepository{DB: db}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	r := gin.Default()

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)

	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/me", authHandler.Me)

	r.Run(":8000")
}
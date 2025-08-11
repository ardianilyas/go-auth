package routes

import (
	"time"

	"github.com/ardianilyas/go-auth/config"
	"github.com/ardianilyas/go-auth/internal/handlers"
	"github.com/ardianilyas/go-auth/internal/middlewares"
	"github.com/ardianilyas/go-auth/internal/models"
	"github.com/ardianilyas/go-auth/internal/repositories"
	"github.com/ardianilyas/go-auth/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() *gin.Engine {
	db, err := gorm.Open(postgres.Open(config.DB_DSN), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	userRepo := &repositories.UserRepository{DB: db}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handler.AuthHandler{AuthService: authService}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/refresh", authHandler.Refresh)
	r.POST("/logout", authHandler.Logout)

	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddleware(), middlewares.CSRFMiddleware())
	auth.GET("/me", authHandler.Me)

	return r
}
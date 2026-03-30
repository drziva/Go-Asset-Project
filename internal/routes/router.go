package routes

import (
	"go-project/internal/handler"
	"go-project/internal/repository"
	"go-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthservice(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	api := r.Group("/api")

	auth := api.Group("/auth")

	auth.POST("/signup", authHandler.SignUp)
	auth.POST("/login", authHandler.Login)

	return r
}

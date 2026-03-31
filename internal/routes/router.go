package routes

import (
	"go-project/internal/config"
	"go-project/internal/handler"
	"go-project/internal/middleware"
	"go-project/internal/repository"
	"go-project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	cfg := config.LoadAuthConfig()
	cookieService := service.NewCookieService(cfg.AccessCookieName, cfg.CookieDomain, cfg.IsProduction)
	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.AccessTokenTTL)

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthservice(userRepo, jwtService)

	authHandler := handler.NewAuthHandler(authService, cookieService, cfg.AccessTokenTTL)

	api := r.Group("/api")

	{
		auth := api.Group("/auth")
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.SignUp)

		auth.Use(middleware.AuthMiddleware(jwtService, cfg.AccessCookieName))
		{
			auth.GET("/me", authHandler.Me)
		}

	}

	return r
}

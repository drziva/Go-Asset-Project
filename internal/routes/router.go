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
	cookieService := service.NewCookieService(cfg.CookieDomain, cfg.IsProduction)
	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.AccessTokenTTL)

	//USER REPO
	userRepo := repository.NewUserRepository(db)

	//MIDDLEWARE
	authMiddleware := middleware.AuthMiddleware(jwtService)
	adminMiddleware := middleware.AdminMiddleware()

	//AUTH
	authService := service.NewAuthservice(userRepo, jwtService)
	authHandler := handler.NewAuthHandler(authService, cookieService, cfg.AccessTokenTTL)

	//ASSETS
	assetRepo := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	api := r.Group("/api")

	{
		auth := api.Group("/auth")
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/logout", authHandler.Logout)

		auth.Use(authMiddleware)
		{
			auth.GET("/me", authHandler.Me)
		}

		protected := api.Group("/")
		protected.Use(authMiddleware)
		{
			//Regular users
			assets := protected.Group("/assets")
			{
				assets.POST("", assetHandler.CreateAsset)
			}

			//Admins
			admin := protected.Group("/admin")
			admin.Use(adminMiddleware)
			{
				adminAssets := admin.Group("/assets")
				{
					adminAssets.GET("", assetHandler.GetAllAssets)
				}
			}
		}

	}

	return r
}

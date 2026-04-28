package routes

import (
	"go-project/internal/client"
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
	googleCfg := client.NewGoogleOauthConfig()

	cookieService := service.NewCookieService(cfg.CookieDomain, cfg.IsProduction)
	jwtService := service.NewJWTService(cfg.JWTSecret, cfg.AccessTokenTTL)

	//USER REPO
	userRepo := repository.NewUserRepository(db)

	//MIDDLEWARE
	authMiddleware := middleware.AuthMiddleware(jwtService)
	adminMiddleware := middleware.AdminMiddleware()

	//AUTH
	authService := service.NewAuthservice(userRepo, googleCfg, jwtService)
	authHandler := handler.NewAuthHandler(authService, cookieService, cfg.AccessTokenTTL)

	//ASSETS
	assetRepo := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	//MICROSERVICE
	microClient := client.NewMicroClient("http://localhost:8081/api")
	microService := service.NewMicroService(microClient)
	microHandler := handler.NewMicroHandler(microService)

	api := r.Group("/api")
	//DEV - MICROSERVICE TEST
	api.GET("/hello", microHandler.GetHello)

	{
		auth := api.Group("/auth")
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/logout", authHandler.Logout)
		auth.GET("/login/google", authHandler.GoogleLogin)

		auth.GET("/google/callback", authHandler.GoogleCallback)
		auth.POST("/google/link-account", authHandler.LinkAndLogin)

		auth.Use(authMiddleware)
		{
			auth.GET("/me", authHandler.Me)
		}

		protected := api.Group("")
		protected.Use(authMiddleware)
		{
			//Regular users
			assets := protected.Group("/assets")
			{
				assets.POST("", assetHandler.CreateAsset)
				assets.GET("", assetHandler.GetAssetsForUser)

				assets.GET("/:id", assetHandler.GetAssetById)
				assets.PUT("/:id", assetHandler.UpdateAsset)
				assets.DELETE("/:id", assetHandler.DeleteAsset)

				assets.GET("/:id/download", assetHandler.DownloadAssetById)

			}

			//Admins
			admin := protected.Group("/admin")
			admin.Use(adminMiddleware)
			{
				adminAssets := admin.Group("/assets")
				{
					adminAssets.GET("", assetHandler.GetAllAssets)
					adminAssets.GET("/:id", assetHandler.GetAnyAssetById)
					adminAssets.PUT("/:id", assetHandler.UpdateAnyAsset)

					adminAssets.GET("/:id/download", assetHandler.DownloadAnyAssetById)
				}
			}
		}

	}

	return r
}

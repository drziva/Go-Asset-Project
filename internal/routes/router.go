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

	//ASSETS
	assetRepo := repository.NewAssetRepository(db)
	assetService := service.NewAssetService(assetRepo)
	assetHandler := handler.NewAssetHandler(assetService)

	//EMAIL MICROSERVICE
	emailClient := client.NewEmailClient("http://localhost:8081/api")
	emailService := service.NewEmailService(emailClient)
	emailHandler := handler.NewEmailHandler(emailService)

	//AUTH
	verificationRepo := repository.NewVerificationCodeRepository(db)
	authService := service.NewAuthservice(userRepo, verificationRepo, googleCfg, jwtService)
	authHandler := handler.NewAuthHandler(authService, emailService, cookieService, cfg.AccessTokenTTL)

	api := r.Group("/api")

	//DEV - EMAIL MICROSERVICE TEST
	email := api.Group("/email")
	email.GET("", emailHandler.SendEmail)
	email.POST("/verification", emailHandler.SendVerificationEmail)

	{
		auth := api.Group("/auth")
		auth.POST("/login", authHandler.Login)
		auth.POST("/signup", authHandler.SignUp)
		auth.POST("/logout", authHandler.Logout)
		auth.GET("/login/google", authHandler.GoogleLogin)

		auth.GET("/google/callback", authHandler.GoogleCallback)
		auth.POST("/google/link-account", authHandler.LinkAndLogin)
		auth.POST("/google/verify-account", authHandler.VerifyLinkAndLogin)

		auth.POST("/forgot-password", authHandler.ForgotPassword)
		auth.POST("/reset-password", authHandler.ResetPassword)

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
					adminAssets.DELETE("/:id", assetHandler.DeleteAnyAsset)

					adminAssets.GET("/:id/download", assetHandler.DownloadAnyAssetById)
				}
			}
		}

	}

	return r
}

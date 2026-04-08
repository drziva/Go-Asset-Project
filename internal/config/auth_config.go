package config

import (
	"go-project/internal/config/utils"
	"os"
)

type AuthConfig struct {
	JWTSecret      string
	AccessTokenTTL int
	IsProduction   bool
	CookieDomain   string
}

func LoadAuthConfig() *AuthConfig {
	secret := utils.GetRequiredEnv("JWT_SECRET")

	//TTL based on prod/dev
	var ttl int
	isProduction := os.Getenv("APP_ENV") == "production"
	if !isProduction {
		ttl = 15 * 24 * 60 * 60 // 15 days
	} else {
		ttl = 15 * 60 //15 minutes
	}

	return &AuthConfig{
		JWTSecret:      secret,
		AccessTokenTTL: ttl, // just for dev
		IsProduction:   isProduction,
		CookieDomain:   os.Getenv("COOKIE_DOMAIN"),
	}
}

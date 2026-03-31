package config

import (
	"log"
	"os"
)

type AuthConfig struct {
	JWTSecret      string
	AccessTokenTTL int
	IsProduction   bool
	CookieDomain   string
}

func LoadAuthConfig() *AuthConfig {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT Secret failed to load")
	}

	var ttl int

	//TTL based on prod/dev
	isProduction := os.Getenv("APP_ENV") == "production"
	if !isProduction {
		ttl = 15 * 24 * 60 * 60
	} else {
		ttl = 15 * 60
	}

	return &AuthConfig{
		JWTSecret:      secret,
		AccessTokenTTL: ttl, // just for dev
		IsProduction:   isProduction,
		CookieDomain:   os.Getenv("COOKIE_DOMAIN"),
	}
}

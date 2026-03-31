package config

import (
	"log"
	"os"
)

type DBConfig struct {
	DSN string
}

func NewDBConfig() *DBConfig {
	DSN := os.Getenv("DSN")
	if DSN == "" {
		log.Fatal("Failed to load DSN")

	}

	return &DBConfig{
		DSN,
	}
}

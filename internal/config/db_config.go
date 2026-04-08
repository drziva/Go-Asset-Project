package config

import (
	"go-project/internal/config/utils"
)

type DBConfig struct {
	DSN string
}

func NewDBConfig() *DBConfig {
	DSN := utils.GetRequiredEnv("DSN")

	return &DBConfig{
		DSN,
	}
}

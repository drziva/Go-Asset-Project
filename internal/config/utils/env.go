package utils

import (
	"log"
	"os"
)

func GetRequiredEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Panicf("Missing required env variable %v", key)
	}

	return value
}

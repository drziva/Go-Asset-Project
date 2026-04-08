package db

import (
	"fmt"
	"go-project/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	cfg := config.NewDBConfig()

	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to Connect with DB: ", err)
	}

	fmt.Println("Connection with DB established!")

	return db
}

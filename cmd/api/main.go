package main

import (
	"go-project/internal/db"
	"go-project/internal/models"
	"go-project/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Failed to load env")
	}

	db := db.NewDatabase()

	db.AutoMigrate(
		&models.User{},
		&models.Asset{},
	)

	r := routes.SetupRoutes(db)

	r.Run()
}

package main

import (
	"go-project/internal/db"
	"go-project/internal/models"
	"go-project/internal/routes"
)

func main() {
	db := db.NewDatabase()

	db.AutoMigrate(&models.User{})

	r := routes.SetupRoutes(db)

	r.Run()
}

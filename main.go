package main

import (
	"log"
	"os"
	"server/src/config"
	"server/src/controller"
	"server/src/repository"
	"server/src/routes"
	"server/src/service"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	config.InitDB()

	clothingRepo := &repository.ClothRepositoryImpl{}

	clothingService := &service.ClothServiceImpl{
		Repo: clothingRepo,
	}

	clothingController := &controller.ClothController{
		Service: clothingService,
	}

	router := routes.RootRoutes(clothingController)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

package main

import (
	"log"

	"github.com/JongSinister/TeeYai_2024/config"
	// "github.com/JongSinister/TeeYai_2024/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Connect to MongoDB
    config.ConnectDB()
    defer config.DisconnectDB()

    // Init Fiber
    app := fiber.New()

    // Set routes
    // routes.SetupRoutes(app)

    // Start server
    log.Fatal(app.Listen(":8080"))
}

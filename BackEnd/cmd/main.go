package main

import (
	"log"
	"os"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)



func main() {

    // Init Fiber
    app := fiber.New()

    if err := godotenv.Load("../config/.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Connect to MongoDB
    config.ConnectDB()
    defer config.DisconnectDB()

    // Initialize database
    config.InitDB()

    // Set routes
    routes.Setup(app)

    // Start server
    port := os.Getenv("PORT")
    log.Fatal(app.Listen("localhost:" + port))
}
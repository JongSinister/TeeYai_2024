package main

import (
	"log"
	"os"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Init Fiber
	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Adjust this to match your frontend origin
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))


    // Load .env file
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

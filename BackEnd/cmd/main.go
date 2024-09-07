package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

// To run backend, use this command in backEnd folder
// CompileDaemon --build="go build -o ./TeeYai.exe main.go" --command=./cmd/TeeYai.exe --directory=./cmd

func main() {

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	envPath := filepath.Join(cwd, "../BackEnd/config/.env")
	log.Printf("Loading .env file from: %s", envPath)

	// Load .env file
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init Fiber
	app := fiber.New()

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // Adjust this to match your frontend origin
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

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

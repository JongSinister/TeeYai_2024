package routes

import (
	"github.com/JongSinister/TeeYai_2024/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    // Set up routes for hospitals and patients
    api := app.Group("/api/v1")

    // Set up routes for authentication
    AuthRoutes(api.Group("/auth"))

    protected := api.Group("/")
    protected.Use(middleware.Protect)
    // Set up routes for orders
    OrderRoutes(api.Group("/orders"))
}
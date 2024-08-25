package routes

import (
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
    // Set up routes for hospitals and patients
    OrderRoutes(app)
}
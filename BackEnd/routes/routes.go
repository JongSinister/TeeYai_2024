package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

    api := app.Group("/api/v1")

    AuthRoutes(api.Group("/auth"))

    OrderRoutes(api.Group("/orders"))
  
}
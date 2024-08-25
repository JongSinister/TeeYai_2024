package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/gofiber/fiber/v3"
)

func OrderRoutes(app *fiber.App) {
    app.Get("/", controllers.GetOrders)
}
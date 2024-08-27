package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	router.Get("/", controllers.GetOrders)
	router.Post("/", controllers.AddOrder)
	router.Get("/:id", controllers.GetOrder)
	router.Delete("/:id", controllers.DeleteOrder)
}
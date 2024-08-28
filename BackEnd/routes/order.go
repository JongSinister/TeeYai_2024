package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/JongSinister/TeeYai_2024/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {
	router.Get("/", controllers.GetOrders, middleware.Protect, middleware.Authorize("admin"))
	router.Post("/", controllers.AddOrder, middleware.Protect)
	router.Get("/:id", controllers.GetOrder, middleware.Protect)
	router.Delete("/:id", controllers.DeleteOrder, middleware.Protect, middleware.Authorize("admin"))
}
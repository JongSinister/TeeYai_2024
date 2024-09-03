package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/JongSinister/TeeYai_2024/middleware"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router fiber.Router) {

	router.Use(middleware.Protect)

	router.Get("/",
		middleware.Protect,
		middleware.Authorize("admin"),
		controllers.GetOrders)

	router.Post("/",
		middleware.Protect,
		controllers.AddOrder)

	router.Get("/:id",
		middleware.Protect,
		controllers.GetOrder)

	router.Delete("/:id",
		middleware.Protect,
		middleware.Authorize("admin"),
		controllers.DeleteOrder)
}

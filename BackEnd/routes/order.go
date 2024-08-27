package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {
	router.Post("/register", controllers.Register)
	router.Post("/login", controllers.Login)
}
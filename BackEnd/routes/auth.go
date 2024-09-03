package routes

import (
	"github.com/JongSinister/TeeYai_2024/controllers"
	"github.com/JongSinister/TeeYai_2024/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(router fiber.Router) {

	router.Post("/register", 
				controllers.Register)

	router.Post("/login", 
				controllers.Login)

	router.Get("/me", 
				middleware.Protect, 
				controllers.Me)

	router.Post("/logout",
				controllers.Logout)

	router.Get("/orders", 
				middleware.Protect, 
				controllers.GetOrdersForUser)
}
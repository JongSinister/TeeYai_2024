package controllers

import (
	"github.com/gofiber/fiber/v3"
)

func GetOrders(c fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "Hello, World!"})
}
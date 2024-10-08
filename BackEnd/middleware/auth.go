package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Protect verifies the JWT token and adds the user to the request context
func Protect(c *fiber.Ctx) error {

	// 1) Get the token from the request header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "No token provided, please Login or Register first"})
	}

	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token"})
	}

	// 2) Extract the claims from the token and add the user to the request context
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token claims"})
	}

	c.Locals("user", claims)
	return c.Next()
}

// Authorize checks if the user has the required role
func Authorize(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// 1) Get the user claims from the request context
		userClaims, ok := c.Locals("user").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User not authorized"})
		}

		// 2) Check if the user has the required role. If so, grant access
		userRole := userClaims["role"].(string)
		for _, role := range roles {
			if userRole == role {
				return c.Next()
			}
		}

		// 3) If the user does not have the required role, return an error
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Access denied"})
	}
}

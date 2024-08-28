package controllers

import (
	"context"
	"os"
	"time"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCollection = "User"


// @desc    Register a new user
// @route   POST /api/v1/auth/register
// @access  Public
func Register(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	// Validate the email format
	if !user.ValidateEmail() {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid email format")
	}

	// Hash the user's password
	if err := user.HashPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error hashing password")
	}

	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	user.UserID = res.InsertedID.(primitive.ObjectID)


	// Generate JWT token
	token, err := user.GenerateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error generating token")
	}

	// Return the user object along with the token
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":  user.Name,
		"token": token,
	})

}

// @desc    Login a user
// @route   POST /api/v1/auth/login
// @access  Public
func Login(c *fiber.Ctx) error {
	var loggedInUser models.User

	if err := c.BodyParser(&loggedInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    // Query for the user by email
    var targetUser models.User
    err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": loggedInUser.Email}).Decode(&targetUser)
    if err != nil {
        return c.Status(fiber.StatusNotFound).SendString("User not found")
    }
	
	// Check if the password is correct
	if !targetUser.CheckPassword(loggedInUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"username":    targetUser.Name,
	})
}

// @desc    Get your user profile
// @route   GET /api/v1/auth/me
// @access  Private
func Me(c *fiber.Ctx) error {
	var user models.User

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Query for the user by id
	err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"_id": user.UserID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return c.Status(fiber.StatusOK).JSON(user)
}


// @desc    Log user out / clear cookie
// @route   GET /api/v1/auth/logout
// @access  Private
func Logout(c *fiber.Ctx) error {

	// Create a cookie object with an expired date to clear it
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true, // Ensure the cookie is HttpOnly
	}

	// Set the cookie with the HttpOnly attribute
	c.Cookie(&cookie)

	// Respond with a success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"date":    time.Now().Format(time.RFC3339),
	})
}



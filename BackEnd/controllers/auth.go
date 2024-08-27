package controllers

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to register user")
	}

	user.UserID = res.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(user)
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
	if targetUser.Password != loggedInUser.Password {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid password")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"username":    targetUser.Name,
	})
}


package controllers

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCollection = "User"


// @desc    Register a new user
// @route   POST /api/v1/auth/register
// @access  Public
func Register(c *fiber.Ctx) error {
	var user models.User
	log.Println("Registering user")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid request format"})
	}

	// Validate the email format
	if !user.ValidateEmail() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid email format"})
	}

	// check if the email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := config.DB.Collection(userCollection).CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Failed to check for existing user"})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message":"Email already exists"})
	}


	// Hash the user's password
	if err := user.HashPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create user"})
	}

	user.UserID = res.InsertedID.(primitive.ObjectID)


	// Generate JWT token
	token, err := user.GenerateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error generating token"})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid request format"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    // Query for the user by email
    var targetUser models.User
    err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": loggedInUser.Email}).Decode(&targetUser)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
    }
	
	// Check if the password is correct
	if !targetUser.CheckPassword(loggedInUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := loggedInUser.GenerateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error generating token"})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"username":    targetUser.Name,
		"token": token,
	})
}

// @desc    Get your user profile
// @route   GET /api/v1/auth/me
// @access  Private
func Me(c *fiber.Ctx) error {

	// ตั้ง user ไว้ สุดท้ายก็ null ดิ แก้ไงวะ
	userEmail, ok := c.Locals("user").(jwt.MapClaims)["email"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error fetching user data"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
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



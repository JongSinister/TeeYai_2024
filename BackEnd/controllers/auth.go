package controllers

import (
	"context"
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

	// 1) Parse the request body into a User struct
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid request format"})
	}

	// 2) Validate the email format
	if !user.ValidateEmail() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid email format"})
	}

	// 3) check if the email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := config.DB.Collection(userCollection).CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Failed to check for existing user"})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message":"Email already exists"})
	}
	///////////////////////////

	// 4) Hash the user's password
	if err := user.HashPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
	}
	///////////////////////////
	
	// 5) Insert the user into the database
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := config.DB.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create user"})
	}

	user.UserID = res.InsertedID.(primitive.ObjectID)
	///////////////////////////

	// 6) Generate JWT token
	token, err := user.GenerateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error generating token"})
	}

	return SendCookie(c, fiber.StatusOK, token, user.UserID)
}

// @desc    Login a user
// @route   POST /api/v1/auth/login
// @access  Public
func Login(c *fiber.Ctx) error {

	// 1) Parse the request body into a User struct
	var loggedInUser models.User

	if err := c.BodyParser(&loggedInUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid request format"})
	}

	// 2) Validate the email format
	if !loggedInUser.ValidateEmail() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Invalid email format"})
	}

	// 3) Find the user in the database by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

    var targetUser models.User
    err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": loggedInUser.Email}).Decode(&targetUser)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
    }
	
	// 4) Check if the password is correct
	if !targetUser.CheckPassword(loggedInUser.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// 5) Generate JWT token
	token, err := targetUser.GenerateToken(os.Getenv("JWT_SECRET"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error generating token"})
	}
	

	return SendCookie(c, fiber.StatusOK, token, targetUser.UserID)
}

// @desc    Get your user profile
// @route   GET /api/v1/auth/me
// @access  Private
func Me(c *fiber.Ctx) error {

	// 1) Get the user's email from the JWT claims
	userEmail, ok := c.Locals("user").(jwt.MapClaims)["email"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error fetching user data"})
	}

	// 2) Find the user in the database by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// @desc    Get user orders
// @route   GET /api/v1/auth/orders
// @access  Private
func GetOrdersForUser(c *fiber.Ctx) error {
	
	// 1) Get the user's email from the JWT claims
	userEmail, ok := c.Locals("user").(jwt.MapClaims)["email"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error fetching user data"})
	}


	// 2) Find the user in the database by email
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"User not found"})
	}

	// 3) For all orders in the user's orders array, find the order in the database
	var orders []map[string]int
	for _, orderID := range user.Orders {
		var order models.Order
		err := config.DB.Collection("Order").FindOne(ctx, bson.M{"_id": orderID}).Decode(&order)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Failed to fetch order"})
		}
		orders = append(orders, order.FoodList)
	}
	
	if len(orders) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"No orders found"})
	}

	return c.JSON(orders)
}



// @desc    Log user out / clear cookie
// @route   GET /api/v1/auth/logout
// @access  Private
func Logout(c *fiber.Ctx) error {

	// 1) Create a cookie object with an expired date to clear it
	cookie := fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true, // Ensure the cookie is HttpOnly
	}

	// 2) Send the cookie to the client
	c.Cookie(&cookie)

	// 3) Respond with a success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"date":    time.Now().Format(time.RFC3339),
	})
}

// send cookie
func SendCookie(c *fiber.Ctx, statusCode int,token string, userID primitive.ObjectID) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"token":   token,
		"userid":  userID,
	})
}

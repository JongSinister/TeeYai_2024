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

const orderCollection = "Order"

//@desc    Get all orders
//@route   GET /api/v1/orders/
//@access  Private
func GetOrders(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB cursor to iterate through the documents
	cursor, err := config.DB.Collection(orderCollection).Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	
	if err := cursor.All(ctx, &orders); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error fetching orders",
		})
	}

	return c.JSON(orders)
}

//@desc    Get order by ID
//@route   GET /api/v1/orders/:id
//@access  Private
func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order := models.Order{}

	err = config.DB.Collection(orderCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(order)
}

//@desc    Add order
//@route   POST /api/v1/orders/
//@access  Private
func AddOrder(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())	
	}

	order.CreatedAt = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := config.DB.Collection(orderCollection).InsertOne(ctx, order)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())	
	}
	 
	order.OrderID = res.InsertedID.(primitive.ObjectID)
	return c.Status(fiber.StatusCreated).JSON(order)
}
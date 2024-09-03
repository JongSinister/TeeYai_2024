package controllers

import (
	"context"
	"time"

	"github.com/JongSinister/TeeYai_2024/config"
	"github.com/JongSinister/TeeYai_2024/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const orderCollection = "Order"

//@desc    Get all orders
//@route   GET /api/v1/orders/
//@access  Private
func GetOrders(c *fiber.Ctx) error {

	// 1) Prepare the query to fetch all orders
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 2) Sort the orders by the created_at field in descending order
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	
	// 3) Fetch all orders from the database
	cursor, err := config.DB.Collection(orderCollection).Find(ctx, bson.M{}, opts)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error fetching orders"})
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	
	// 4) Iterate over the cursor and decode each order
	if err := cursor.All(ctx, &orders); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Error fetching orders",
			"msg": err,
		})
	}

	if len(orders) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No orders found"})
	}

	return c.JSON(orders)
}

//@desc    Get order by ID
//@route   GET /api/v1/orders/:id
//@access  Private
func GetOrder(c *fiber.Ctx) error {

	// 1) Get the order ID from the URL and convert it to an ObjectID
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID format"})
	}

	// 2) Fetch the order from the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	order := models.Order{}

	err = config.DB.Collection(orderCollection).FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error fetching order"})
	}

	return c.JSON(order)
}


//@desc    Add order
//@route   POST /api/v1/orders/
//@access  Private
func AddOrder(c *fiber.Ctx) error {

	// 1) Get the user's email from the JWT claims
	userEmail, ok := c.Locals("user").(jwt.MapClaims)["email"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Error fetching user data"})
	}

	// 1) Fetch the user from the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := config.DB.Collection(userCollection).FindOne(ctx, bson.M{"email": userEmail}).Decode(&user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error fetching user data"})
	}

	// 2) Parse the request body into a Order struct
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request format"})	
	}

	// 3) get the user ID from the user object, add created_at timestamp to the order
	order.UserID = user.UserID
	order.CreatedAt = primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond))

	// 4) Insert the order into the database
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := config.DB.Collection(orderCollection).InsertOne(ctx, order)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error creating order"})
	}
	 
	order.OrderID = res.InsertedID.(primitive.ObjectID)

	// 5) After the order is created, add the order ID to the user's orders array
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = config.DB.Collection(userCollection).UpdateOne(ctx, bson.M{"_id": order.UserID}, bson.M{"$push": bson.M{"orders": order.OrderID}})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error updating user orders"})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}


//@desc    Delete order by ID
//@route   DELETE /api/v1/orders/:id
//@access  Private
func DeleteOrder(c *fiber.Ctx) error {

	// 1) Get the order ID from the URL and convert it to an ObjectID
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID format"})
	}

	// 2) Delete the order from the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := config.DB.Collection(orderCollection).DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error deleting order"})
	}
	
	if res.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "error"})
	}

	// 3) After the order is deleted, remove the order ID from the user's orders array
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	_, err = config.DB.Collection(userCollection).UpdateOne(ctx, bson.M{"orders": objectID}, bson.M{"$pull": bson.M{"orders": objectID}})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error updating user orders"})
	}

	

	return c.JSON(fiber.Map{"message": "Order deleted successfully"})
}
// package controllers

// import (
// 	"github.com/JongSinister/TeeYai_2024/config"
// 	"github.com/JongSinister/TeeYai_2024/models"
//     "context"
//     "log"
//     "time"

//     "github.com/gofiber/fiber/v3"
//     "go.mongodb.org/mongo-driver/bson"
//     "go.mongodb.org/mongo-driver/mongo"
// )

// func CreateUser(c *fiber.Ctx) error {
//     user := new(models.User)

//     if err := c.BodyParser(user); err != nil {
//         return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
//     }

//     collection := config.DB.Database("your_database_name").Collection("users")

//     ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     res, err := collection.InsertOne(ctx, bson.D{{"name", user.Name}, {"email", user.Email}})
//     if err != nil {
//         log.Fatal(err)
//         return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
//     }

//     return c.Status(fiber.StatusCreated).JSON(fiber.Map{"insertedId": res.InsertedID})
// }

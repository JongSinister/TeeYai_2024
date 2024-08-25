package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
    OrderID   primitive.ObjectID       `bson:"_id,omitempty"`          
    UserID    string                   `bson:"userid" validate:"required"`       
    FoodList  map[string]int           `bson:"foodList"`     
    CreatedAt primitive.DateTime       `bson:"createdAt" validate:"required"`   
}

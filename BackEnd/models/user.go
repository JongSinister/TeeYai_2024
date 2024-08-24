package models

import (
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Name      string             `bson:"name" validate:"required"`
    Email     string             `bson:"email" validate:"required"`
    Role      string             `bson:"role,omitempty"`
    Password  string             `bson:"password" validate:"required"`
    CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
}

// Check Email Validation
func (h *User) ValidateEmail() bool {
    regex := `^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
    re := regexp.MustCompile(regex)
    return re.MatchString(h.Email)
}
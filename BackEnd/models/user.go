package models

import (
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    UserID    primitive.ObjectID    `bson:"_id,omitempty"`
    Name      string                `bson:"name" validate:"required"`
    Email     string                `bson:"email" validate:"required"`
    Role      string                `bson:"role,omitempty"`
    Orders    []primitive.ObjectID  `bson:"orders,omitempty"`
    Password  string                `bson:"password" validate:"required"`
    CreatedAt primitive.DateTime    `bson:"created_at,omitempty"`
}

// Check Email Validation
func (h *User) ValidateEmail() bool {
    regex := `^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
    re := regexp.MustCompile(regex)
    return re.MatchString(h.Email)
}

// HashPassword hashes the user's password using bcrypt
func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword compares a hashed password with the provided password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// GenerateToken generates a JWT token for the user
func (u *User) GenerateToken(secret string) (string, error) {
	claims := jwt.MapClaims{
		"id":   u.UserID,
		"role": u.Role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
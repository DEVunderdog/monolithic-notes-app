package models

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id"`
	First_name *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name  *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password   *string            `json:"password" validate:"required,min=6,passwd"`
	Email      *string            `json:"email" validate:"email,required"`
	Phone      *string            `json:"phone" validate:"required"`
	Token      *string            `json:"token"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	User_id    string             `json:"user_id"`
	// User_type  *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`

}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasMinLength := len(password) >= 6
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(password)

	return hasMinLength && hasUppercase && hasLowercase && hasSpecial
}



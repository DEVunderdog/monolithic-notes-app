package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       *string            `json:"title" bson:"title"`
	Description *string            `json:"description" bson:"description"`
	Created_at  time.Time          `json:"created_at" bson:"created_at"`
	Updated_at  time.Time          `json:"updated_at" bson:"updated_at"`
	User_id     string             `json:"user_id" bson:"user_id"`
	Notes_id    string             `json:"notes_id" bson:"notes_id"`
}

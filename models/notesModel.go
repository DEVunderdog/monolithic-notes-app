package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       *string            `json:"title"`
	Description *string            `json:"description"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
	Notes_id    string             `json:"notes_id"`
}

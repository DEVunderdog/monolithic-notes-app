package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DEVunderdog/monolithic-notes-app/database"
	"github.com/DEVunderdog/monolithic-notes-app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var notesCollection *mongo.Collection = database.OpenNotesCollection(database.Client)

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var notes models.Notes

		if err := c.BindJSON(&notes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		notes.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		notes.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		notes.ID = primitive.NewObjectID()
		notes.Notes_id = notes.ID.Hex()

		resultInsertionNumber, insertErr := notesCollection.InsertOne(ctx, notes)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}
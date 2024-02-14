package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DEVunderdog/monolithic-notes-app/database"
	"github.com/DEVunderdog/monolithic-notes-app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var notesCollection *mongo.Collection = database.OpenNotesCollection(database.Client)

func CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var notes models.Notes


		uidValue, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve uid"})
			return
		}

		uid, ok := uidValue.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "uid not a string"})
			return
		}

		if err := c.BindJSON(&notes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		notes.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		notes.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		notes.ID = primitive.NewObjectID()
		notes.Notes_id = notes.ID.Hex()
		notes.User_id = uid
		

		resultInsertionNumber, insertErr := notesCollection.InsertOne(ctx, notes)
		if insertErr != nil {
			msg := fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		c.JSON(http.StatusOK, resultInsertionNumber)
	}

}

func GetNotes() gin.HandlerFunc{
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var notes []models.Notes

		uidValue, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "uid not a string"})
			return
		}

		uid, ok := uidValue.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "uid not a string"})
			return
		}

		cursor, err := notesCollection.Find(ctx, bson.M{"user_id": uid})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while querying for the notes " + err.Error()})
			return
		}

		if err = cursor.All(ctx, &notes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while decoding for the cursor: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"notes": notes})
	}
}

func UpdateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		
		noteId := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(noteId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		}

		filter := bson.M{"_id": objectID}

		var note models.Notes

		if err := c.BindJSON(&note); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		updateObj := bson.M{
			"$set": bson.M{
				"title": note.Title,
				"description": note.Description,
				"updated_at": time.Now(),
			},
		}

		_, err = notesCollection.UpdateOne(ctx, filter, updateObj)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating note"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully"})
	}
}

func DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		noteId := c.Param("id")
		objectID, err := primitive.ObjectIDFromHex(noteId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note id"})
		}

		filter := bson.M{"_id": objectID}

		_, err = notesCollection.DeleteOne(ctx, filter)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating note"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Note Deleted successfully"})

		
	}
}
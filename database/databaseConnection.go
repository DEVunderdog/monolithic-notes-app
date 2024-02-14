package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	MongoDb := os.Getenv("MONGODB_URL")
	optionMongoDbUrl := MongoDb + "?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(optionMongoDbUrl)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

var Client *mongo.Client = DBInstance()

func CreateUniqueIndexes(collection *mongo.Collection) {
	emailIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), emailIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	phoneIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "phone", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context.Background(), phoneIndexModel)
	if err != nil {
		log.Fatal(err)
	}
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	CreateUniqueIndexes(collection)
	return collection
}

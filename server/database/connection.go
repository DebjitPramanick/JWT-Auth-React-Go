package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTION_URL = "mongodb+srv://debjit:Gyb6XLHxpxvHsRgn@cluster0.c4dqrmw.mongodb.net/?retryWrites=true&w=majority"
const DB_NAME = "jwtAuth"
const COLLECTION_NAME = "users"

var Collection *mongo.Collection

func Connect() {
	clientOption := options.Client().ApplyURI(CONNECTION_URL)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to DB")

	Collection = client.Database(DB_NAME).Collection(COLLECTION_NAME)

	fmt.Println("Collection instance is ready.")
}
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected!")

	database := client.Database("testdb")
	collection := database.Collection("test")

	filter := bson.M{"title": "test"}

	var result bson.M
	fmt.Println("filter:", filter)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query result:", result)
}

// MongoDB Connected!
// filter: map[title:test]
// Query result: map[_id:ObjectID("64a61b8c8f48f12d1a485487") title:test]

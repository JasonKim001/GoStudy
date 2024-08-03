package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("testdb").Collection("testcollection")

	// Insert a single document
	doc := bson.D{{"name", "Alice"}, {"age", 30}}
	_, err = collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Document inserted!")

	// Query the document
	var result bson.D
	err = collection.FindOne(context.TODO(), bson.D{{"name", "Alice"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %v\n", result)
}

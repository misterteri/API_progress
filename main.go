package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

type Students []Student

const mongoURL string = "mongodb://localhost:27017"

func connectMongo(url string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return client
}

func main() {
	mongoClient := connectMongo(mongoURL)
	defer mongoClient.Disconnect(context.TODO())
	fmt.Println("Connected to MongoDB")

	collection := mongoClient.Database("test").Collection("students")

	var student = Students{
		{
			Name:   "Gilbert",
			Age:    21,
			Gender: "Male",
		},
		{
			Name:   "Bella",
			Age:    20,
			Gender: "Female",
		},
		{
			Name:   "Jett",
			Age:    20,
			Gender: "Male",
		},
		{
			Name:   "Jett",
			Age:    20,
			Gender: "Male",
		},
	}

	for _, student := range student {
		filter := bson.M{"name": student.Name}
		var result Student
		err := collection.FindOne(context.TODO(), filter).Decode(&result) // ini buat find data
		if err != mongo.ErrNoDocuments {
			fmt.Println("Found a result: ", result)
		} else {
			_, err := collection.InsertOne(context.TODO(), student) // ini buat insert data
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

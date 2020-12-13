package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type trainer struct {
	Name string
	Age  int
	City string
}

func insertTrainer(x trainer, y *mongo.Collection) string {
	_, err := y.InsertOne(context.TODO(), x)
	if err != nil {
		log.Fatal(err)
		return "Unsucess insert"
	}
	return "Sucess insert"

}
func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:1234")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection := client.Database("test").Collection("trainer")

	ash := trainer{"Ash", 10, "Hatyai"}
	t := insertTrainer(ash, collection)
	fmt.Println(t)
	filter := bson.D{{"name", "Ash"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Update Error ")
	} else {
		fmt.Println("Update Success")
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")
}

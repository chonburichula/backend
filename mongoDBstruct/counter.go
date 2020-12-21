package mongoDBstruct

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type counter struct {
	ID            string `bson:"_id"`
	SequenceValue int    `bson:"sequence_value"`
}

func connectToCounterCollection() (*mongo.Client, *mongo.Collection, error) {
	var collection *mongo.Collection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return client, collection, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return client, collection, err
	}
	collection = client.Database("chulachonburi").Collection("counter")
	return client, collection, err
}

func CreateNewCounter(sequenceName string) (*mongo.InsertOneResult, error) {
	var insertResult *mongo.InsertOneResult
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return insertResult, err
	}
	counter := counter{
		ID:            sequenceName,
		SequenceValue: 0,
	}
	collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: sequenceName}})
	insertResult, err = collection.InsertOne(context.TODO(), counter)
	if err != nil {
		return insertResult, err
	}
	err = disConnectToDatbase(client)
	return insertResult, err
}

func GetNextApplicantID() int {
	client, collection, _ := connectToApplicantCollection()
	filter := bson.D{{Key: "_id", Value: "id"}}
	//option := options.FindOneAndUpdate()
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "sequence_value", Value: 1}}}}
	var result counter
	_ = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	_ = disConnectToDatbase(client)
	return result.SequenceValue
}

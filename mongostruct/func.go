package mongostruct

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToApplicantCollection() (*mongo.Client, *mongo.Collection, error) {
	var collection *mongo.Collection
	clientOptions := options.Client().ApplyURI("mongodb://54.255.211.157:27017")
	clientOptions.SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, err
	}
	collection = client.Database("chulachonburi").Collection("applicant")
	return client, collection, err
}

func connectToCounterCollection() (*mongo.Client, *mongo.Collection, error) {
	var collection *mongo.Collection
	clientOptions := options.Client().ApplyURI("mongodb://54.255.211.157:27017")
	clientOptions.SetServerSelectionTimeout(5 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, nil, err
	}
	collection = client.Database("chulachonburi").Collection("counter")
	return client, collection, err
}

func disConnectToDatbase(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	return err
}

//CreateNewCounter is ...
func createNewCounter(sequenceName string) (*mongo.InsertOneResult, error) {
	var insertResult *mongo.InsertOneResult
	client, collection, err := connectToCounterCollection()
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

func getNextApplicantID() int {
	client, collection, _ := connectToCounterCollection()
	filter := bson.D{{Key: "_id", Value: "id"}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "sequence_value", Value: 1}}}}

	var result counter
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	if err != nil {
		createNewCounter("id")
		collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&result)
	}
	_ = disConnectToDatbase(client)
	return result.SequenceValue + 1
}

//Insert is ...
func Insert(applicant Applicant) (*mongo.InsertOneResult, error) {
	applicant.ID = getNextApplicantID()
	applicant.Graded = false
	applicant.Score = 0
	var insertResult *mongo.InsertOneResult
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return nil, err
	}
	insertResult, err = collection.InsertOne(context.TODO(), applicant)
	if err != nil {
		return nil, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return nil, err
	}
	return insertResult, err
}

//UpdateGraded is ...
func UpdateStatusAndScore(applicantID int, newScore int32) (*mongo.UpdateResult, error) {
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: applicantID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: true}, {Key: "score", Value: newScore}}}}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func SearchApplicantByName(searchName string) ([]Applicant, error) {
	var applicants []Applicant
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "name", Value: searchName}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp Applicant
		err = cur.Decode(&temp)
		if err != nil {
			return nil, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}

func GetGradedApplicant() ([]ApplicantOnlyScore, error) {
	var applicants []ApplicantOnlyScore
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "status", Value: true}}
	option := options.Find()
	option.SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "score", Value: 1}})
	option.SetSort(bson.D{{Key: "score", Value: -1}})
	cur, err := collection.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp ApplicantOnlyScore
		err = cur.Decode(&temp)
		if err != nil {
			return nil, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}

func GetUnGradedApplicant() ([]ApplicantOnlyScore, error) {
	var applicants []ApplicantOnlyScore
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return nil, err
	}
	filter := bson.D{{Key: "status", Value: false}}
	option := options.Find()
	option.SetProjection(bson.D{{Key: "_id", Value: 1}, {Key: "score", Value: 1}})
	cur, err := collection.Find(context.TODO(), filter, option)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp ApplicantOnlyScore
		err = cur.Decode(&temp)
		if err != nil {
			return nil, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return nil, err
	}
	return applicants, nil
}

func GetApplicantAnswer(applicantID int) (ApplicantOnlyAnswer, error) {
	var applicant ApplicantOnlyAnswer
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return applicant, err
	}
	filter := bson.D{{Key: "_id", Value: applicantID}}
	option := options.FindOne().SetProjection(bson.D{
		{Key: "_id", Value: 1},
		{Key: "gradinganswer1", Value: 1},
		{Key: "gradinganswer2", Value: 1},
		{Key: "gradinganswer3", Value: 1},
		{Key: "gradinganswer4", Value: 1},
		{Key: "answer1", Value: 1},
		{Key: "answer2", Value: 1},
		{Key: "answer3", Value: 1},
		{Key: "answer4", Value: 1},
		{Key: "answer5", Value: 1},
		{Key: "answer6", Value: 1},
		{Key: "answer7", Value: 1},
	})
	err = collection.FindOne(context.TODO(), filter, option).Decode(&applicant)
	if err != nil {
		return applicant, err
	}
	err = disConnectToDatbase(client)
	if err != nil {
		return applicant, err
	}
	return applicant, nil
}

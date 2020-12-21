package mongoDBstruct

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Applicant struct {
	ID                int32  `bson:"id" json:"id" binding:"gte=1,"`
	Email             string `bson:"email" json:"email" validate:"email,required"`
	Title             string `bson:"title" json:"title" binding:"required"`
	Name              string `bson:"name" json:"name"`
	Surname           string `bson:"surname" json:"surname"`
	Nickname          string `bson:"nickname" json:"nickname"`
	Birthdate         string `bson:"birthdate" json:"birthdate"`
	Age               string `bson:"age" json:"age"`
	BloodType         string `bson:"bloodtype" json:"bloodtype"`
	Religion          string `bson:"religion" json:"religion"`
	Address           string `bson:"address" json:"address"`
	PhoneNumber       string `bson:"phonenumber" json:"phonenumber"`
	LineID            string `bson:"lineid" json:"lineid"`
	Facebook          string `bson:"facebook" json:"facebook"`
	Class             string `bson:"class" json:"class"`
	Major             string `bson:"major" json:"major"`
	School            string `bson:"school" json:"school"`
	Disease           string `bson:"disease" json:"disease"`
	Medicine          string `bson:"medicine" json:"medicine"`
	FoodLimitation    string `bson:"foodlimitation" json:"foodlimitation"`
	ClothSize         string `bson:"clothsize" json:"clothsize"`
	FatherName        string `bson:"fathername" json:"fathername"`
	FatherPhoneNumber string `bson:"fatherphonenumber" json:"fatherphonenumber"`
	MotherName        string `bson:"mothername" json:"mothername"`
	MotherPhoneNumber string `bson:"motherphonenumber" json:"motherphonenumber"`
	ParentName        string `bson:"parentname" json:"parentname"`
	ParentType        string `bson:"parenttype" json:"parenttype"`
	ParentPhoneNumber string `bson:"parentphonenumber" json:"parentphonenumber"`
	GradingAnswer1    string `bson:"gradinganswer1" json:"gradinganswer1"`
	GradingAnswer2    string `bson:"gradinganswer2" json:"gradinganswer2"`
	GradingAnswer3    string `bson:"gradinganswer3" json:"gradinganswer3"`
	Answer1           string `bson:"answer1" json:"answer1"`
	Answer2           string `bson:"answer2" json:"answer2"`
	Answer3           string `bson:"answer3" json:"answer3"`
	Answer4           string `bson:"answer4" json:"answer4"`
	Answer5           string `bson:"answer5" json:"answer5"`
	Answer6           string `bson:"answer6" json:"answer6"`
	Answer7           string `bson:"answer7" json:"answer7"`
	Status            string `bson:"status" json:"status" `
	Score             int    `bson:"score" json:"score"`
}

func connectToApplicantCollection() (*mongo.Client, *mongo.Collection, error) {
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
	collection = client.Database("chulachonburi").Collection("applicant")
	return client, collection, err
}

func disConnectToDatbase(client *mongo.Client) error {
	err := client.Disconnect(context.TODO())
	return err
}

func Insert(applicant Applicant, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	var insertResult *mongo.InsertOneResult
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return insertResult, err
	}
	insertResult, err = collection.InsertOne(context.TODO(), applicant)
	if err != nil {
		return insertResult, err
	}
	err = disConnectToDatbase(client)
	return insertResult, err
}

func UpdateGraded(applicant Applicant) (*mongo.UpdateResult, error) {
	var updateResult *mongo.UpdateResult
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return updateResult, err
	}
	filter := bson.D{{Key: "id", Value: applicant.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "graded"}}}}
	updateResult, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return updateResult, err
	}
	err = disConnectToDatbase(client)
	return updateResult, err
}

func ShowUnGradedApplicant() ([]Applicant, error) {
	var applicants []Applicant
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return applicants, err
	}
	filter := bson.D{{Key: "status", Value: "ungraded"}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return applicants, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp Applicant
		err = cur.Decode(&temp)
		if err != nil {
			return applicants, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return applicants, err
	}
	err = disConnectToDatbase(client)
	return applicants, err
}

func ShowGradedApplicant() ([]Applicant, error) {
	var applicants []Applicant
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return applicants, err
	}
	filter := bson.D{{Key: "status", Value: "graded"}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return applicants, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp Applicant
		err = cur.Decode(&temp)
		if err != nil {
			return applicants, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return applicants, err
	}
	err = disConnectToDatbase(client)
	return applicants, err
}

func SearchApplicantByName(searchName string) ([]Applicant, error) {
	var applicants []Applicant
	client, collection, err := connectToApplicantCollection()
	if err != nil {
		return applicants, err
	}
	filter := bson.D{{Key: "name", Value: searchName}}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return applicants, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var temp Applicant
		err = cur.Decode(&temp)
		if err != nil {
			return applicants, err
		}
		applicants = append(applicants, temp)
	}
	err = cur.Err()
	if err != nil {
		return applicants, err
	}
	err = disConnectToDatbase(client)
	return applicants, err
}

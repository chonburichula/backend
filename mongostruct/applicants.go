package mongostruct

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Applicant struct {
	ID                int32  `bson:"_id" json:"_id" `
	Email             string `bson:"email" json:"email" validate:"email,required"`
	Title             string `bson:"title" json:"title" binding:"required"`
	Name              string `bson:"name" json:"name" binding:"required"`
	Surname           string `bson:"surname" json:"surname" binding:"required"`
	Nickname          string `bson:"nickname" json:"nickname" binding:"required"`
	Birthdate         string `bson:"birthdate" json:"birthdate" binding:"required"`
	Age               int32  `bson:"age" json:"age" binding:"required,gte=1"`
	BloodType         string `bson:"bloodtype" json:"bloodtype" binding:`
	Religion          string `bson:"religion" json:"religion" binding:"required"`
	Address           string `bson:"address" json:"address" binding:"required"`
	PhoneNumber       string `bson:"phonenumber" json:"phonenumber" binding:"required"`
	LineID            string `bson:"lineid" json:"lineid" binding:"required"`
	Facebook          string `bson:"facebook" json:"facebook" binding:"required"`
	Class             string `bson:"class" json:"class" binding:"required"`
	Major             string `bson:"major" json:"major" binding:"required"`
	School            string `bson:"school" json:"school" binding:"required"`
	Disease           string `bson:"disease" json:"disease" binding:"required"`
	Medicine          string `bson:"medicine" json:"medicine" binding:"required"`
	FoodLimitation    string `bson:"foodlimitation" json:"foodlimitation" binding:"required"`
	ClothSize         string `bson:"clothsize" json:"clothsize" binding:"required"`
	FatherName        string `bson:"fathername" json:"fathername" binding:"required"`
	FatherPhoneNumber string `bson:"fatherphonenumber" json:"fatherphonenumber" binding:"required"`
	MotherName        string `bson:"mothername" json:"mothername" binding:"required"`
	MotherPhoneNumber string `bson:"motherphonenumber" json:"motherphonenumber" binding:"required"`
	ParentName        string `bson:"parentname" json:"parentname" `
	ParentType        string `bson:"parenttype" json:"parenttype"`
	ParentPhoneNumber string `bson:"parentphonenumber" json:"parentphonenumber"`
	GradingAnswer1    string `bson:"gradinganswer1" json:"gradinganswer1" binding:"required"`
	GradingAnswer2    string `bson:"gradinganswer2" json:"gradinganswer2" binding:"required"`
	GradingAnswer3    string `bson:"gradinganswer3" json:"gradinganswer3" binding:"required"`
	Answer1           string `bson:"answer1" json:"answer1" binding:"required"`
	Answer2           string `bson:"answer2" json:"answer2" binding:"required"`
	Answer3           string `bson:"answer3" json:"answer3" binding:"required"`
	Answer4           string `bson:"answer4" json:"answer4" binding:"required"`
	Answer5           string `bson:"answer5" json:"answer5" binding:"required"`
	Answer6           string `bson:"answer6" json:"answer6" binding:"required"`
	Answer7           string `bson:"answer7" json:"answer7" binding:"required"`
	Status            string `bson:"status" json:"status" binding:"requoired, `
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

//Insert is ...
func Insert(applicant Applicant) (*mongo.InsertOneResult, error) {
	applicant.ID = GetNextApplicantID()
	applicant.Status = "ungraded"
	applicant.Score = 0
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

//UpdateGraded is ...
func (applicant Applicant) UpdateGraded() (*mongo.UpdateResult, error) {
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

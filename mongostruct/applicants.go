package mongostruct

type Applicant struct {
	ID                int    `bson:"_id" json:"_id"`
	Email             string `bson:"email" json:"email" `
	Title             string `bson:"title" json:"title" `
	Name              string `bson:"name" json:"name" `
	Surname           string `bson:"surname" json:"surname" `
	Nickname          string `bson:"nickname" json:"nickname" `
	Birthdate         string `bson:"birthdate" json:"birthdate" `
	BloodType         string `bson:"bloodtype" json:"bloodtype" `
	Religion          string `bson:"religion" json:"religion" `
	Address           string `bson:"address" json:"address" `
	PhoneNumber       string `bson:"phonenumber" json:"phonenumber" `
	LineID            string `bson:"lineid" json:"lineid" `
	Facebook          string `bson:"facebook" json:"facebook" `
	Class             string `bson:"class" json:"class"`
	Major             string `bson:"major" json:"major"`
	School            string `bson:"school" json:"school"`
	Disease           string `bson:"disease" json:"disease"`
	Medicine          string `bson:"medicine" json:"medicine"`
	FoodLimitation    string `bson:"foodlimitation" json:"foodlimitation"`
	ClothSize         string `bson:"clothsize" json:"clothsize"`
	ParentName        string `bson:"parentname" json:"parentname" `
	ParentType        string `bson:"parenttype" json:"parenttype"`
	ParentPhoneNumber string `bson:"parentphonenumber" json:"parentphonenumber"`
	GradingAnswer1    string `bson:"gradinganswer1" json:"gradinganswer1"`
	GradingAnswer2    string `bson:"gradinganswer2" json:"gradinganswer2"`
	GradingAnswer3    string `bson:"gradinganswer3" json:"gradinganswer3"`
	GradingAnswer4    string `bson:"gradinganswer4" json:"gradinganswer4"`
	Answer1           string `bson:"answer1" json:"answer1"`
	Answer2           string `bson:"answer2" json:"answer2"`
	Answer3           string `bson:"answer3" json:"answer3"`
	Answer4           string `bson:"answer4" json:"answer4"`
	Answer5           string `bson:"answer5" json:"answer5"`
	Answer6           string `bson:"answer6" json:"answer6"`
	Answer7           string `bson:"answer7" json:"answer7"`
	Graded            bool   `bson:"graded" json:"graded"`
	Score             int    `bson:"score" json:"score"`
}

type ApplicantOnlyScore struct {
	ID    int `bson:"_id" json:"_id"`
	Score int `bson:"score" json:"score"`
}

type ApplicantOnlyAnswer struct {
	ID             int    `bson:"_id" json:"_id"`
	GradingAnswer1 string `bson:"gradinganswer1" json:"gradinganswer1"`
	GradingAnswer2 string `bson:"gradinganswer2" json:"gradinganswer2"`
	GradingAnswer3 string `bson:"gradinganswer3" json:"gradinganswer3"`
	GradingAnswer4 string `bson:"gradinganswer4" json:"gradinganswer4"`
	Answer1        string `bson:"answer1" json:"answer1"`
	Answer2        string `bson:"answer2" json:"answer2"`
	Answer3        string `bson:"answer3" json:"answer3"`
	Answer4        string `bson:"answer4" json:"answer4"`
	Answer5        string `bson:"answer5" json:"answer5"`
	Answer6        string `bson:"answer6" json:"answer6"`
	Answer7        string `bson:"answer7" json:"answer7"`
}

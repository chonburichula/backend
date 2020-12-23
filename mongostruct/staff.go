package mongostruct

type staff struct {
	username string `bson:"username" json:"username"`
	password string `bson:"password" json:"password"`
}

package model

type User struct {
	UserId   string `bson:"user_id",omitempty"`
	username string `bson:"username"`
	password string `bson:"password"`
}

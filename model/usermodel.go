package model

type User struct {
	UserId   string `bson:"user_id",omitempty"`
	Username string `bson:"username"`
	password string `bson:"password"`
}

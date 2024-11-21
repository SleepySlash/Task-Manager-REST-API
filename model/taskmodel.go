package model

type Task struct {
	UserId string `bson:"user_id,omitempty"`
	Name   string `bson:"name"`
	Date   string `bson:"created_on"`
	Status string `bson:"status"`
}

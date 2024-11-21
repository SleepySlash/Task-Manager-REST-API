package model

type Task struct {
	UserId    string `bson:"user_id,omitempty"`
	name      string `bson:"name"`
	createdOn string `bson:"created_on"`
	status    bool   `bson:"status"`
}

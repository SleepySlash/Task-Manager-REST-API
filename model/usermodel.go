package model

type User struct {
	UserId   string `json:"user_id,omitempty"`
	Username string `json:"username"`
	password string `json:"password"`
}

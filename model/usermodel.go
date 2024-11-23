package model

type User struct {
	UserId   string `json:"userid,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

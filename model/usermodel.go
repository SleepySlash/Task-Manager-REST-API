package model

type User struct {
	username string `json:"username,omitempty"`
	password string `json:"password,omitempty"`
}

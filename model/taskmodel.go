package model

type Task struct {
	UserId string `json:"userid,omitempty"`
	Name   string `json:"name"`
	Date   string `json:"created_on"`
	Status string `json:"status"`
}

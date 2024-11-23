package model

type Task struct {
	UserId string `json:"userid,omitempty"`
	Name   string `json:"name"`
	Date   string `json:"created_on"`
	status string `json:"status"`
}

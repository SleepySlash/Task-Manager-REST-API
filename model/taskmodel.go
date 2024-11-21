package model

type Task struct {
	name      string `json:"name,omitempty"`
	createdOn string `json:"created_on,omitempty"`
	status    bool   `json:"status,omitempty"`
}

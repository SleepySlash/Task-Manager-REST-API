package services

import (
	"Task-Manager-REST-API/model"
)

type TaskService interface {
	NewTask(name string, userid string) (model.Task, error)
	GetTask(userid string, name string, date string) (model.Task, error)
	GetAllTasks(userid string) ([]model.Task, error)
	DeleteTask(userid string, name string, date string) (model.Task, error)
	DeleteAllTasks(userid string) error
	UpdateTask(userid string, newTask string, name string, date string) (model.Task, error)
}
type taskService struct {
	repo model.Tasks
}

func NewTaskService(repo model.Tasks) TaskService {
	return &taskService{
		repo: repo,
	}
}
func (t *taskService) NewTask(name string, userid string) (model.Task, error) {
}
func (t *taskService) GetTask(userid string, name string, date string) (model.Task, error)    {}
func (t *taskService) GetAllTasks(userid string) ([]model.Task, error)                        {}
func (t *taskService) DeleteTask(userid string, name string, date string) (model.Task, error) {}
func (t *taskService) DeleteAllTasks(userid string) error                                     {}
func (t *taskService) UpdateTask(userid string, newTask string, name string, date string) (model.Task, error) {
}

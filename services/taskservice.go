package services

import "Task-Manager-REST-API/model"

type TaskService interface {
	newTask(name string) (model.Task, error)
	getTask(name string) (model.Task, error)
	getAllTasks() ([]model.Task, error)
	deleteTask(name string) (model.Task, error)
	deleteAllTasks() error
	updateTask(name string) (model.Task, error)
}
type taskService struct {
	repo model.Tasks
}

func NewTaskService(repo model.Tasks) TaskService {
	return &taskService{
		repo: repo,
	}
}
func (t *taskService) newTask(name string) (model.Task, error)    {}
func (t *taskService) getTask(name string) (model.Task, error)    {}
func (t *taskService) getAllTasks() ([]model.Task, error)         {}
func (t *taskService) deleteTask(name string) (model.Task, error) {}
func (t *taskService) deleteAllTasks() error                      {}
func (t *taskService) updateTask(name string) (model.Task, error) {}

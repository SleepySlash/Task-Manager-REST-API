package services

import (
	"Task-Manager-REST-API/model"
	"log"
	"time"
)

type TaskService interface {
	NewTask(name string, userid string) (model.Task, error)
	GetTask(userid string, name string, date string) (model.Task, error)
	GetAllTasks(userid string) ([]model.Task, error)
	DeleteTask(userid string, name string, date string) (model.Task, error)
	DeleteAllTasks(userid string) error
	UpdateTask(userid string, newTask string, name string, date string) (model.Task, error)
	MarkDone(userid string, name string, date string) (model.Task, error)
	MarkUnDone(userid string, name string, date string) (model.Task, error)
}
type taskService struct {
	repo model.Tasks
}

func NewTaskService(repo model.Tasks) TaskService {
	return &taskService{
		repo: repo,
	}
}

// Service Layer Function for creating a new task
func (t *taskService) NewTask(name string, userid string) (model.Task, error) {
	createdtime := time.Now()
	formattedDate := createdtime.Format("2006-01-02")
	task := model.Task{
		UserId: userid,
		Name:   name,
		Date:   formattedDate,
	}
	err := t.repo.Create(task)
	if err != nil {
		log.Println("error during creation of task")
		return task, err
	}
	return task, nil
}

// Service Layer Function for getting an existing task
func (t *taskService) GetTask(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Get(userid, name, date)
	if err != nil {
		log.Println("error while searching for the task")
		return result, err
	}
	return result, nil
}

// Service Layer Function for getting all existing task of the user
func (t *taskService) GetAllTasks(userid string) ([]model.Task, error) {
	var result []model.Task
	result, err := t.repo.All(userid)
	if err != nil {
		log.Println("error while searching for all the tasks of the user")
		return result, err
	}
	return result, nil
}

// Service Layer Function for deleting an existing task
func (t *taskService) DeleteTask(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Delete(userid, name, date)
	if err != nil {
		log.Println("error while searching for the task")
		return result, err
	}
	return result, nil
}

// Service Layer Function for deleting all existing task of the user
func (t *taskService) DeleteAllTasks(userid string) error {
	err := t.repo.DeleteAll(userid)
	if err != nil {
		log.Println("error while deleting all the tasks of the user")
		return err
	}
	return nil
}

// Service Layer Function for updating an existing task
func (t *taskService) UpdateTask(userid string, newTask string, name string, date string) (model.Task, error) {
	updatedTask := model.Task{
		UserId: userid,
		Name:   newTask,
		Date:   date,
	}

	result, err := t.repo.Update(updatedTask, name)
	if err != nil || result < 1 {
		log.Println("error while updating the task")
		return updatedTask, err
	}
	return updatedTask, nil
}
func (t *taskService) MarkDone(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Done(userid, name, date)
	if err != nil {
		log.Println("error while marking the task done")
		return result, err
	}
	return result, nil
}

func (t *taskService) MarkUnDone(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Undone(userid, name, date)
	if err != nil {
		log.Println("error while marking the task undone")
		return result, err
	}
	return result, nil
}

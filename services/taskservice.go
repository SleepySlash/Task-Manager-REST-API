package services

import (
	"Task-Manager-REST-API/model"
	"log"
	"time"
)

type TaskService interface {
	NewTask(name string, userid string) (model.Task, error)
	NewTasks(name []string, userid string) ([]model.Task, error)
	GetTask(userid string, name string, date string) (model.Task, error)
	GetAllTasks(userid string) ([]model.Task, error)
	GetAllDoneInclusive(userid string) ([]model.Task, error)
	DeleteTask(userid string, name string, date string) (model.Task, error)
	DeleteAllTasks(userid string) (string, error)
	UpdateTask(userid string, newTask string, name string, date string) (model.Task, string)
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
	// date format for the tasks is yyyy-mm-dd
	createdtime := time.Now()
	formattedDate := createdtime.Format("2006-01-02")
	task := model.Task{
		UserId: userid,
		Name:   name,
		Date:   formattedDate,
		Status: "pending",
	}
	err := t.repo.Create(task)
	if err != nil {
		log.Println("error during creation of task")
		return task, err
	}
	return task, nil
}

// Service Layer Function for creating a multiple new tasks
func (t *taskService) NewTasks(names []string, userid string) ([]model.Task, error) {
	// date format for the tasks is yyyy-mm-dd
	createdtime := time.Now()
	formattedDate := createdtime.Format(time.Stamp)
	var tasks []model.Task
	for _, taskname := range names {
		var task model.Task

		task.UserId = userid
		task.Name = taskname
		task.Date = formattedDate
		task.Status = "pending"

		tasks = append(tasks, task)
	}
	err := t.repo.CreateMany(tasks)
	if err != nil {
		log.Println("error during creation of mulitple tasks")
		return tasks, err
	}
	return tasks, nil
}

// Service Layer Function for getting an existing task
func (t *taskService) GetTask(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	log.Println(userid, name, date)
	result, err := t.repo.Get(userid, name, date)
	if err != nil {
		log.Println("error while searching for the task")
		return result, err
	}
	return result, nil
}

// Service Layer Function for getting all the tasks of the user which are pending
func (t *taskService) GetAllTasks(userid string) ([]model.Task, error) {
	var result []model.Task // the tasks array which will be returned to the user
	result, err := t.repo.All(userid)
	if err != nil {
		log.Println("error while searching for all the tasks of the user")
		return result, err
	}
	return result, nil
}

// Service Layer Function for getting all existing task of the user
func (t *taskService) GetAllDoneInclusive(userid string) ([]model.Task, error) {
	var result []model.Task // the tasks array which will be returned to the user
	result, err := t.repo.AllTasks(userid)
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
func (t *taskService) DeleteAllTasks(userid string) (string, error) {
	count, err := t.repo.DeleteAll(userid)
	if err != nil {
		log.Println("error while deleting all the tasks of the user")
		return "", err
	}
	if count > 1 {
		return "delted all the users", nil
	}
	return "no users found to be deleted", nil
}

// Service Layer Function for updating an existing task
func (t *taskService) UpdateTask(userid string, newTask string, name string, date string) (model.Task, string) {
	updatedTask := model.Task{
		UserId: userid,
		Name:   newTask,
		Date:   date,
		Status: "pending",
	}

	result, err := t.repo.Update(updatedTask, name)
	if err != nil || result < 1 {
		log.Println("error while updating the task")
		return updatedTask, "error in update"
	}
	return updatedTask, ""
}

// Service Layer Function for marking a task as done
func (t *taskService) MarkDone(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Done(userid, name, date)
	if err != nil {
		log.Println("error while marking the task done")
		return result, err
	}
	return result, nil
}

// Service Layer Function for marking a task as pending
func (t *taskService) MarkUnDone(userid string, name string, date string) (model.Task, error) {
	var result model.Task
	result, err := t.repo.Undone(userid, name, date)
	if err != nil {
		log.Println("error while marking the task undone")
		return result, err
	}
	return result, nil
}

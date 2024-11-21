package controllers

import (
	"Task-Manager-REST-API/model"
	"Task-Manager-REST-API/services"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type TaskController interface {
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	GetTheTask(w http.ResponseWriter, r *http.Request)
	GetAllTheTasks(w http.ResponseWriter, r *http.Request)
	UpdateTheTask(w http.ResponseWriter, r *http.Request)
	DeleteTheTask(w http.ResponseWriter, r *http.Request)
	DeleteAllTheTasks(w http.ResponseWriter, r *http.Request)
}
type taskcontroller struct {
	service services.TaskService
}

func Tasker(client *mongo.Client) TaskController {
	db := client.Database(os.Getenv("TASK_DB")).Collection(os.Getenv("TASK_COLLECTION"))
	collection := model.CreateTaskRepo(db)

	repo := services.NewTaskService(collection)
	return &taskcontroller{
		service: repo,
	}
}

func (c *taskcontroller) CreateNewTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) GetTheTask(w http.ResponseWriter, r *http.Request)        {}
func (c *taskcontroller) GetAllTheTasks(w http.ResponseWriter, r *http.Request)    {}
func (c *taskcontroller) UpdateTheTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) DeleteTheTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) DeleteAllTheTasks(w http.ResponseWriter, r *http.Request) {}

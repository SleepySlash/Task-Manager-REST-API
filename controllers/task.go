package controllers

import (
	"Task-Manager-REST-API/middleware"
	"Task-Manager-REST-API/model"
	"Task-Manager-REST-API/services"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskController interface {
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	GetTheTask(w http.ResponseWriter, r *http.Request)
	GetAllTheTasks(w http.ResponseWriter, r *http.Request)
	GetAllIncludingDone(w http.ResponseWriter, r *http.Request)
	UpdateTheTask(w http.ResponseWriter, r *http.Request)
	DeleteTheTask(w http.ResponseWriter, r *http.Request)
	DeleteAllTheTasks(w http.ResponseWriter, r *http.Request)
	MarkTheTaskComplete(w http.ResponseWriter, r *http.Request)
	MarkTheTaskPending(w http.ResponseWriter, r *http.Request)
}
type taskcontroller struct {
	service services.TaskService
}

func Tasker(client *mongo.Client) TaskController {
	db := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("TASK_COLLECTION"))
	collection := model.CreateTaskRepo(db)

	repo := services.NewTaskService(collection)
	return &taskcontroller{
		service: repo,
	}
}

// Handler function to create a new task
func (c *taskcontroller) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	taskName := r.FormValue("task")
	log.Println("creating a todo for ", taskName)

	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	result, err := c.service.NewTask(taskName, userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Entered new task"))
	json.NewEncoder(w).Encode(result)
}

// Handler function to get a specific task
func (c *taskcontroller) GetTheTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]
	log.Println(name, date)
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	result, err := c.service.GetTask(userid, name, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("fetched the task"))
	json.NewEncoder(w).Encode(result)
}

// Handler function to Get all the tasks of the user which are pending
func (c *taskcontroller) GetAllTheTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	result, err := c.service.GetAllTasks(userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("found all the tasks"))
	json.NewEncoder(w).Encode(result)
}

// Handler function to Get all the tasks of the user
func (c *taskcontroller) GetAllIncludingDone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userid, err := middleware.GetIdFromContext(r.Context())
	log.Print(userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	result, err := c.service.GetAllDoneInclusive(userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("found all the tasks"))
	json.NewEncoder(w).Encode(result)
}

// Handler function to update a specific task
func (c *taskcontroller) UpdateTheTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]
	newTask := r.FormValue("task")
	log.Println("updating a todo,", name, " to ", newTask)
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	result, er := c.service.UpdateTask(userid, newTask, name, date)
	if er != "" {
		log.Println(er)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("updated the task"))
	json.NewEncoder(w).Encode(result)

}

// Handler function to delete a specific task
func (c *taskcontroller) DeleteTheTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	result, err := c.service.DeleteTask(userid, name, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("deleted the task"))
	json.NewEncoder(w).Encode(result)
}

// Handler function to delete all the tasks of the user
func (c *taskcontroller) DeleteAllTheTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	err = c.service.DeleteAllTasks(userid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("deleted all the tasks"))
}

func (c *taskcontroller) MarkTheTaskComplete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	result, err := c.service.MarkDone(userid, name, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("marked the task complete"))
	json.NewEncoder(w).Encode(result)
}

func (c *taskcontroller) MarkTheTaskPending(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	date := mux.Vars(r)["date"]
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	result, err := c.service.MarkUnDone(userid, name, date)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("marked the task pending"))
	json.NewEncoder(w).Encode(result)
}

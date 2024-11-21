package controllers

import (
	"Task-Manager-REST-API/services"
	"net/http"
)

type TaskController interface {
	createNewTask(w http.ResponseWriter, r *http.Request)
	getTheTask(w http.ResponseWriter, r *http.Request)
	getAllTheTasks(w http.ResponseWriter, r *http.Request)
	updateTheTask(w http.ResponseWriter, r *http.Request)
	deleteTheTask(w http.ResponseWriter, r *http.Request)
	deleteAllTheTasks(w http.ResponseWriter, r *http.Request)
}
type taskcontroller struct {
	service services.TaskService
}

// func NewTaskService() TaskController{
// 	&taskcontroller{
// 		service: // code,
// 	}
// }

func (c *taskcontroller) createNewTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) getTheTask(w http.ResponseWriter, r *http.Request)        {}
func (c *taskcontroller) getAllTheTasks(w http.ResponseWriter, r *http.Request)    {}
func (c *taskcontroller) updateTheTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) deleteTheTask(w http.ResponseWriter, r *http.Request)     {}
func (c *taskcontroller) deleteAllTheTasks(w http.ResponseWriter, r *http.Request) {}

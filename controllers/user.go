package controllers

import (
	"Task-Manager-REST-API/services"
	"net/http"
)

type UserController interface {
	createNewUser(w http.ResponseWriter, r *http.Request)
	getTheUser(w http.ResponseWriter, r *http.Request)
	getAllTheUsers(w http.ResponseWriter, r *http.Request)
	updateTheUser(w http.ResponseWriter, r *http.Request)
	deleteTheUser(w http.ResponseWriter, r *http.Request)
	deleteAllTheUsers(w http.ResponseWriter, r *http.Request)
}
type usercontroller struct {
	service services.UserService
}

// func NewUserService() UserController{
// 	&usercontroller{
// 		service: // code,
// 	}
// }

func (c *usercontroller) createNewUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) getTheUser(w http.ResponseWriter, r *http.Request)        {}
func (c *usercontroller) getAllTheUsers(w http.ResponseWriter, r *http.Request)    {}
func (c *usercontroller) updateTheUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) deleteTheUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) deleteAllTheUsers(w http.ResponseWriter, r *http.Request) {}

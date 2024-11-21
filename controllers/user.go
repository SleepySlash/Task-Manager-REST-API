package controllers

import (
	"Task-Manager-REST-API/model"
	"Task-Manager-REST-API/services"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
	GetTheUser(w http.ResponseWriter, r *http.Request)
	GetAllTheUsers(w http.ResponseWriter, r *http.Request)
	UpdateTheUser(w http.ResponseWriter, r *http.Request)
	DeleteTheUser(w http.ResponseWriter, r *http.Request)
	DeleteAllTheUsers(w http.ResponseWriter, r *http.Request)
}
type usercontroller struct {
	service services.UserService
}

func User(client *mongo.Client) UserController {
	db := client.Database(os.Getenv("USER_DB")).Collection(os.Getenv("USER_COLLECTION"))
	collection := model.CreateUserRepo(db)

	repo := services.NewUserService(collection)
	return &usercontroller{
		service: repo,
	}
}

func (c *usercontroller) CreateNewUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	result, err := c.service.NewUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("registered a new user"))
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
func (c *usercontroller) GetTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	ok, err := c.service.GetUser(user)
	if !ok {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("user logged in successfully"))
	w.WriteHeader(http.StatusAccepted)
}
func (c *usercontroller) GetAllTheUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	result, err := c.service.GetAllUser()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("registered a new user"))
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
func (c *usercontroller) UpdateTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	result, err := c.service.UpdateUser(name, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("updated the user info"))
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
func (c *usercontroller) DeleteTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	name := mux.Vars(r)["name"]
	result, err := c.service.DeleteUser(name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("deleted the user"))
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(result)
}
func (c *usercontroller) DeleteAllTheUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	err := c.service.DeleteAllUser()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("deleted all the users"))
	w.WriteHeader(http.StatusAccepted)
}

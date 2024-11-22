package controllers

import (
	"Task-Manager-REST-API/middleware"
	"Task-Manager-REST-API/model"
	"Task-Manager-REST-API/services"
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	db := client.Database(os.Getenv("DATABASE")).Collection(os.Getenv("USER_COLLECTION"))
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
	err := c.service.NewUser(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("registered a new user"))
	w.WriteHeader(http.StatusAccepted)
}
func (c *usercontroller) GetTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	token, err := c.service.GetUser(user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token)
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
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	ok, err := c.service.UpdateUser(userid, user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte("updated the user info"))
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(ok)
}
func (c *usercontroller) DeleteTheUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	userid, err := middleware.GetIdFromContext(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusExpectationFailed)
	}
	result, err := c.service.DeleteUser(userid)
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

package controllers

import (
	"Task-Manager-REST-API/model"
	"Task-Manager-REST-API/services"
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
	db := client.Database(os.Getenv("USER_DB")).Collection(os.Getenv("USER_COLLECTION"))
	collection := model.CreateUserRepo(db)

	repo := services.NewUserService(collection)
	return &usercontroller{
		service: repo,
	}
}

func (c *usercontroller) CreateNewUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) GetTheUser(w http.ResponseWriter, r *http.Request)        {}
func (c *usercontroller) GetAllTheUsers(w http.ResponseWriter, r *http.Request)    {}
func (c *usercontroller) UpdateTheUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) DeleteTheUser(w http.ResponseWriter, r *http.Request)     {}
func (c *usercontroller) DeleteAllTheUsers(w http.ResponseWriter, r *http.Request) {}

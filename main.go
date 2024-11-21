package main

import (
	"Task-Manager-REST-API/controllers"
	"Task-Manager-REST-API/middleware"
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	client := middleware.DataSource()
	defer client.Disconnect(context.Background())
	users := controllers.User(client)
	tasks := controllers.Tasker(client)

	todo := router.PathPrefix("/todo").Subrouter()

	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/login", users.GetTheUser).Methods("POST")
	todo.HandleFunc("/getall", tasks.GetAllTheTasks).Methods("GET")

	todo.Use(middleware.AuthMiddleware)

	http.ListenAndServe(":3000", router)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Manager is running on port 3000"))
}

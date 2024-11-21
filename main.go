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

	user := router.PathPrefix("/user").Subrouter()
	todo := router.PathPrefix("/todo").Subrouter()

	users := controllers.User(client)
	tasks := controllers.Tasker(client)

	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/login", users.GetTheUser).Methods("POST")
	router.HandleFunc("/register", users.GetTheUser).Methods("POST")
	router.HandleFunc("/delete", users.DeleteAllTheUsers).Methods("DELETE")

	user.HandleFunc("/update", users.UpdateTheUser).Methods("PUT")
	user.HandleFunc("/delete", users.DeleteTheUser).Methods("DELETE")

	todo.HandleFunc("/new", tasks.CreateNewTask).Methods("POST")
	todo.HandleFunc("/get", tasks.GetTheTask).Methods("POST")
	todo.HandleFunc("/getall", tasks.GetAllTheTasks).Methods("GET")
	todo.HandleFunc("/update", tasks.UpdateTheTask).Methods("PUT")
	todo.HandleFunc("/delete", tasks.DeleteTheTask).Methods("DELETE")
	todo.HandleFunc("/deleteall", tasks.DeleteAllTheTasks).Methods("DELETE")
	todo.HandleFunc("/done", tasks.MarkTheTaskComplete).Methods("DELETE")
	todo.HandleFunc("/pending", tasks.MarkTheTaskPending).Methods("DELETE")

	todo.Use(middleware.AuthMiddleware)
	user.Use(middleware.AuthMiddleware)

	http.ListenAndServe(":3000", router)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Manager is running on port 3000"))
}

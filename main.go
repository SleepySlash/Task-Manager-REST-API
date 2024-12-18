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

	// health check no need of auth
	router.HandleFunc("/health", health).Methods("GET")
	// login handler no need of auth
	router.HandleFunc("/login", users.GetTheUser).Methods("POST")
	// register handler no need of auth
	router.HandleFunc("/register", users.CreateNewUser).Methods("POST")
	// delete all user handler no need of auth
	router.HandleFunc("/delete", users.DeleteAllTheUsers).Methods("DELETE")

	// delete all user handler auth required
	user.HandleFunc("/update", users.UpdateTheUser).Methods("PUT")
	// delete the user handler auth required
	// user.HandleFunc("/delete", users.DeleteTheUser).Methods("DELETE")

	// task related handlers, auth required
	todo.HandleFunc("/new", tasks.CreateNewTask).Methods("POST")
	todo.HandleFunc("/newtasks", tasks.CreateNewTasks).Methods("POST")
	todo.HandleFunc("/get/{name}/{date}", tasks.GetTheTask).Methods("GET")
	todo.HandleFunc("/gettasks", tasks.GetAllTheTasks).Methods("GET")
	todo.HandleFunc("/getall", tasks.GetAllIncludingDone).Methods("GET")
	todo.HandleFunc("/update/{name}/{date}", tasks.UpdateTheTask).Methods("PUT")
	todo.HandleFunc("/mark/done", tasks.MarkGivenTasksComplete).Methods("PUT")
	todo.HandleFunc("/mark/done/{name}/{date}", tasks.MarkTheTaskComplete).Methods("PUT")
	todo.HandleFunc("/mark/pending/{name}/{date}", tasks.MarkTheTaskPending).Methods("PUT")
	todo.HandleFunc("/delete/{name}/{date}", tasks.DeleteTheTask).Methods("DELETE")
	todo.HandleFunc("/deleteall", tasks.DeleteAllTheTasks).Methods("DELETE")

	router.Use(middleware.RequestLogger)
	todo.Use(middleware.AuthMiddleware)
	user.Use(middleware.AuthMiddleware)

	http.ListenAndServe(":3000", router)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Manager is running on port 3000"))
}

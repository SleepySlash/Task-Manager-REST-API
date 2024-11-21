package main

import (
	"Task-Manager-REST-API/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	todo := router.PathPrefix("/todo").Subrouter()

	router.HandleFunc("/health", health).Methods("GET")

	todo.Use(middleware.AuthMiddleware)

	http.ListenAndServe(":3000", router)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Manager is running on port 3000"))
}

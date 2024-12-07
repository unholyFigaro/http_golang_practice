package main

import (
	"net/http"

	"github.com/gorilla/mux"
	db "github.com/unxly/golang-pa/internal/database"
	"github.com/unxly/golang-pa/internal/handlers"
	"github.com/unxly/golang-pa/internal/taskService"
)

var task string = "User"

func main() {
	DB := db.InitDB()
	DB.AutoMigrate(&taskService.Task{})
	repo := taskService.NewTaskRepository(DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.New(service)

	router := mux.NewRouter()
	router.HandleFunc("/task", handler.PostTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/task", handler.GetTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/task/{id:[0-9]+}", handler.UpdateTaskHandler).Methods(http.MethodPatch)
	router.HandleFunc("/task/{id:[0-9]+}", handler.DeleteTaskHandler).Methods(http.MethodDelete)
	http.ListenAndServe("localhost:9090", router)
}

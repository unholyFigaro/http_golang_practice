package main

import (
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/unxly/golang-pa/models"
	"github.com/unxly/golang-pa/pkg/db"
	"github.com/unxly/golang-pa/pkg/handlers"
)

var task string = "User"

func main() {
	DB := db.InitDB()
	h := handlers.New(DB)
	DB.AutoMigrate(&m.Message{})
	// fmt.Println(DB)
	router := mux.NewRouter()
	router.HandleFunc("/task", h.SetTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("/task", h.GreetingTask).Methods(http.MethodGet)
	http.ListenAndServe("localhost:9090", router)
}

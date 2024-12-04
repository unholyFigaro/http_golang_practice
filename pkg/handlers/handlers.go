package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "github.com/unxly/golang-pa/models"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{DB: db}
}

func (h handler) SetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task m.Message
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	// task = message.Message
	log.Printf("Adding task: %+v", task)
	res := h.DB.Create(&task)
	if res.Error != nil {
		http.Error(w, "Something went wrong, try again later", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task was updated"))
}

func (h handler) GreetingTask(w http.ResponseWriter, r *http.Request) {
	var tasks []m.Message
	t := h.DB.Find(&tasks)
	fmt.Printf("%v", tasks)
	if t.Error != nil {
		http.Error(w, "Something went wrong, try again later", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

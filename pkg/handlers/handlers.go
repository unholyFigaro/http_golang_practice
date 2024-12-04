package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
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
	res := h.DB.Create(&task)
	if res.Error != nil {
		http.Error(w, "Something went wrong, try again later", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task was added"))
}

func (h handler) GreetingTask(w http.ResponseWriter, r *http.Request) {
	var tasks []m.Message
	t := h.DB.Find(&tasks)
	if t.Error != nil {
		http.Error(w, "Something went wrong, try again later", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func (h handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var task m.Message
	id := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var existingTask m.Message
	if err = h.DB.First(&existingTask, id).Error; err != nil {
		http.Error(w, "Task with id not found", http.StatusNotFound)
		return
	}
	if err = h.DB.Model(&existingTask).Updates(task).Error; err != nil {
		http.Error(w, "Something went wrong, try again later", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task was updated"))
}

func (h handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task m.Message
	if err := h.DB.Model(&task).Delete(&task, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
}

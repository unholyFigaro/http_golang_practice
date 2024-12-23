package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unxly/golang-pa/internal/taskService"
	"github.com/unxly/golang-pa/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func New(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PostTasks(w http.ResponseWriter, r *http.Request) {
	var task taskService.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}
	// task = message.Message
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

func (h *Handler) GetTasks(
	_ context.Context,
	_ tasks.GetTasksRequestObject,
) (tasks.GetTasksResponseObject, error) {
	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasks200JSONResponse

	for _, task := range tasks {

	}
}

func (h *Handler) UpdateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var task taskService.Task

	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	id := uint(idUint)

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := h.Service.UpdateTaskByID(id, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

func (h *Handler) DeleteTasksId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	idUint, err := strconv.ParseInt(idStr, 10, 32)
	id := uint(idUint)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	err = h.Service.DeleteTaskByID(id)
	w.WriteHeader(http.StatusOK)
}

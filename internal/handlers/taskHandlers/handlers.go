package taskHandlers

import (
	"context"

	"github.com/unxly/golang-pa/internal/taskService"
	"github.com/unxly/golang-pa/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func New(service *taskService.TaskService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PostTasks(_ context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := req.Body
	userIdFromReq := uint(*req.Body.UserId)
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserId: *&userIdFromReq,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	return response, nil
}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	tsks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, v := range tsks {
		task := tasks.Task{
			Id:     &v.ID,
			Task:   &v.Task,
			IsDone: &v.IsDone,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) GetUsersIdTasks(_ context.Context, request tasks.GetUsersIdTasksRequestObject) (tasks.GetUsersIdTasksResponseObject, error) {
	userId := request.Id
	tsks, err := h.Service.GetAllTasksForUser(uint(userId))
	if err != nil {
		return nil, err
	}
	response := tasks.GetUsersIdTasks200JSONResponse{}
	for _, v := range tsks {
		task := tasks.Task{
			Id:     &v.ID,
			Task:   &v.Task,
			IsDone: &v.IsDone,
			UserId: &v.UserId,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	IdToPatch := request.Id
	requestBody := request.Body
	taskToUpdate := taskService.Task{
		Task:   *requestBody.Task,
		IsDone: *requestBody.IsDone,
	}
	updatedTask, err := h.Service.UpdateTaskByID(IdToPatch, taskToUpdate)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}
	return response, nil
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	IdToDelete := request.Id
	deletedTask, err := h.Service.DeleteTaskByID(IdToDelete)
	if err != nil {
		return nil, err
	}
	response := tasks.DeleteTasksId200JSONResponse{
		Id:     &deletedTask.ID,
		Task:   &deletedTask.Task,
		IsDone: &deletedTask.IsDone,
	}
	return response, nil
}

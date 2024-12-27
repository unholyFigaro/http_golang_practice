package userHandlers

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/unxly/golang-pa/internal/userService"
	"github.com/unxly/golang-pa/internal/web/users"
)

type Handler struct {
	Service *userService.UserService
}

func New(service *userService.UserService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (r *Handler) PostUsers(_ context.Context, req users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	reqBody := req.Body
	userToCreate := userService.User{
		Password: *reqBody.Password,
		Email:    *reqBody.Email,
	}
	log.Printf("handler: %v", userToCreate)
	log.Printf("handler: %v", reqBody)
	createdUser, err := r.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	response := users.PostUsers201JSONResponse{
		Id:       &createdUser.Id,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return response, nil
}

func (r *Handler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := r.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var respUsers users.GetUsers200JSONResponse
	for _, v := range allUsers {
		tempUser := users.User{
			Id:    &v.Id,
			Email: &v.Email,
		}
		respUsers = append(respUsers, tempUser)
	}
	return respUsers, nil
}

func (r *Handler) PatchUsersId(_ context.Context, req users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	idToPatch := req.Id
	reqBody := req.Body
	userToUpdate := userService.User{
		Email:    *reqBody.Email,
		Password: *reqBody.Password,
	}
	updatedUser, err := r.Service.UpdateUserById(idToPatch, userToUpdate)
	if err != nil {
		return nil, err
	}
	response := users.PatchUsersId200JSONResponse{
		Id:    &updatedUser.Id,
		Email: &updatedUser.Email,
	}
	return response, nil
}

func (r *Handler) DeleteUsersId(_ context.Context, req users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	idToDelete := req.Id
	deletedUser, err := r.Service.DeleteUserById(idToDelete)
	if err != nil {
		return nil, err
	}
	response := users.DeleteUsersId200JSONResponse{
		Id:    &deletedUser.Id,
		Email: &deletedUser.Email,
	}
	return response, nil
}

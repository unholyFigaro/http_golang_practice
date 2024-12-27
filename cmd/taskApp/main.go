package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/unxly/golang-pa/internal/database"
	"github.com/unxly/golang-pa/internal/handlers/taskHandlers"
	"github.com/unxly/golang-pa/internal/handlers/userHandlers"
	"github.com/unxly/golang-pa/internal/taskService"
	"github.com/unxly/golang-pa/internal/userService"
	"github.com/unxly/golang-pa/internal/web/tasks"
	"github.com/unxly/golang-pa/internal/web/users"
)

func main() {
	db.InitDB()
	log.Println(db.DB)
	err := db.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		return
	}
	err = db.DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatalf("failed to migrate User: %v", err)
	}
	taskRepo := taskService.NewTaskRepository(db.DB)
	taskService := taskService.NewTaskService(taskRepo)

	userRepo := userService.NewUserRepository(db.DB)
	userServiceInstance := userService.NewUserService(userRepo)

	taskHandler := taskHandlers.New(taskService)
	userHandler := userHandlers.New(userServiceInstance)
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	userStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, userStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

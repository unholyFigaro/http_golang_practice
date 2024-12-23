package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/unxly/golang-pa/internal/database"
	"github.com/unxly/golang-pa/internal/handlers"
	"github.com/unxly/golang-pa/internal/taskService"
	"github.com/unxly/golang-pa/internal/web/tasks"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(db.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.New(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/unxly/golang-pa/internal/database"
	"github.com/unxly/golang-pa/internal/handlers/taskHandlers"
	"github.com/unxly/golang-pa/internal/taskService"
	"github.com/unxly/golang-pa/internal/web/tasks"
)

func main() {
	db.InitDB()
	log.Println(db.DB)
	err := db.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		return
	}

	taskRepo := taskService.NewTaskRepository(db.DB)
	taskService := taskService.NewTaskService(taskRepo)

	taskHandler := taskHandlers.New(taskService)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	taskStrictHandler := tasks.NewStrictHandler(taskHandler, nil)
	tasks.RegisterHandlers(e, taskStrictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}

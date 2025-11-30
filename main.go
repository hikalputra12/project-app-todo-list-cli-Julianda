package main

import (
	"todo-list-app/cmd"
	"todo-list-app/service"
)

func main() {
	// Initialize the service
	taskService := service.NewTaskService()
	// Initialize the handler with the service
	taskHandler := cmd.NewTaskHandler(taskService)
	cmd.Execute(taskHandler)
}

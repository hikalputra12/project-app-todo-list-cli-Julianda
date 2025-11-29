package cmd

import (
	"fmt"
	"strings"
	"todo-list-app/model"

	"github.com/jedib0t/go-pretty/v6/table"
)

// TaskServiceInterface interface used by CLI handler
type TaskServiceInterface interface {
	GetAllTask() ([]model.Task, error)
	CreateTask(input model.Task) (model.Task, error)
	GetTaskByTitle(input string) (*model.Task, error) //using pointer for efficiency and capability to return the nil
	DeleteTask(id int) error
	UpdateTask(id int, task model.Task) (model.Task, error)
}
type TaskHandler struct {
	Service TaskServiceInterface
}

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{Service: service}
}

// ListAllTask returns all tasks
func (h *TaskHandler) ListAllTask() (string, error) {
	tasks, err := h.Service.GetAllTask()
	if err != nil {
		return "", err
	}
	if len(tasks) == 0 {
		return "No Task found.\n", nil
	}
	var b strings.Builder //joinning string like using + but with memory efficiency

	//tablewriter receive io.writer as an output
	t := table.NewWriter()
	t.SetOutputMirror(&b) //set output to strings.Builder //masih salah sepertinya
	//set coloumn header
	t.AppendHeader(table.Row{"ID", "Task", "Status", "Priority"})

	//add data
	for _, ts := range tasks {
		t.AppendRow(table.Row{ts.ID, ts.Title, ts.Status, ts.Priority})
	}
	//render table
	t.Render()
	//retrieves and returns the final result of all the data that has been collected into strings.Builder
	return b.String(), nil
}

// CreateTask handle creating a new task and return message
func (h *TaskHandler) CreateTask(title, status, priority string) (string, error) {
	input := model.Task{
		Title:    title,
		Status:   status,
		Priority: priority,
	}

	created, err := h.Service.CreateTask(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task created: ID=%d, Title=%s, Status=%s, Priority=%s\n",
		created.ID, created.Title, created.Status, created.Priority), nil

}

// function to update task
func (h *TaskHandler) UpdateTask(id int, status string) (string, error) {
	input := model.Task{
		Status: status,
	}
	update, err := h.Service.UpdateTask(id, input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Task Update: ID=%d, Title=%s, Status=%s, Priority=%s\n",
		update.ID, update.Title, update.Status, update.Priority), nil

}

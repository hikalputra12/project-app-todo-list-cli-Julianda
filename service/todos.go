package service

import (
	"todo-list-app/model"
	"todo-list-app/utils"
)

// This interface will define whose contract the handler will use.
type TaskServiceInterface interface {
	GetAllTask() ([]model.Task, error)
	CreateTask(input model.Task) (model.Task, error)
	GetTaskByTitle(input string) (*model.Task, error) //using pointer for efficiency and capability to return the nil
	DeleteTask(id int) error
	UpdateTask(id int, task model.Task) (model.Task, error)
}

type TaskService struct{}

// constructor
func NewTaskService() *TaskService {
	return &TaskService{}
}

// function can access data
func (t *TaskService) accessTask() ([]model.Task, error) {
	return utils.ReadFileJson()
}

// function for saving task
func (t *TaskService) saveTask(task []model.Task) error {
	return utils.WriteFileJson(task)
}

// function to get list of task
func (t *TaskService) GetAllTask() ([]model.Task, error) {
	return t.accessTask()
}

// function to get task by Title
func (t *TaskService) GetTaskByTitle(input string) (*model.Task, error) {
	task, err := t.accessTask()
	if err != nil {
		return nil, err
	}
	//access list of task
	for _, ts := range task {
		if ts.Title == input {
			copy := ts
			return &copy, nil
		}
	}
	return nil, utils.ErrNotFOund
}

// function to create a task with an ID
func (t *TaskService) CreateTask(input model.Task) (model.Task, error) {
	tasks, err := t.accessTask()
	if err != nil {
		return model.Task{}, err
	}
	maxID := 0
	for _, ts := range tasks {
		if ts.ID > maxID {
			maxID = ts.ID
		}
	}
	input.ID = maxID + 1

	tasks = append(tasks, input)
	//save input task
	if err := t.saveTask(tasks); err != nil {
		return model.Task{}, err
	}
	return input, nil
}

// functio to delete a task
func (t *TaskService) DeleteTask(id int) error {
	tasks, err := t.accessTask()
	if err != nil {
		return err
	}
	//to initialize new []model.Task chunks efficiently
	newTaskList := make([]model.Task, 0, len(tasks))
	found := false
	for _, ts := range tasks {
		if ts.ID == id {
			found = true
			continue
		}
		newTaskList = append(newTaskList, ts)
	}
	if !found {
		return utils.ErrNotFOund
	}
	return t.saveTask(newTaskList)
}

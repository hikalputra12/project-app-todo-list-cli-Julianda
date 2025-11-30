package service

import (
	"errors"
	"strings"
	"todo-list-app/model"
	"todo-list-app/utils"
)

// This interface will define whose contract the handler will use.
type TaskServiceInterface interface {
	GetAllTask() ([]model.Task, error)
	CreateTask(input model.Task) (model.Task, error)
	GetTaskByTitle(input string) (*model.Task, error) //using pointer for efficiency and capability to return the nil
	DeleteTask(number int) error
	UpdateTask(number int, input model.Task) (model.Task, error)
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

	//validation input status
	InputStatus := strings.ToLower(input.Status)
	if input.Status != "" {
		isValid := false
		if InputStatus == utils.NoProgressStatus ||
			InputStatus == utils.OnProgressStatus ||
			InputStatus == utils.CompletedStatus {
			isValid = true
		}
		if !isValid {
			return model.Task{}, errors.New("status is not valid")
		}
	}

	//validation input priority
	InputPriority := strings.ToLower(input.Priority)
	if InputPriority != "" {
		isValid := false
		if InputPriority == utils.LowPriority ||
			InputPriority == utils.MediumPriority ||
			InputPriority == utils.HighPriority {
			isValid = true
		}
		if !isValid {
			return model.Task{}, utils.ErrNotFOund
		}
	}

	tasks = append(tasks, input)
	//save input task
	if err := t.saveTask(tasks); err != nil {
		return model.Task{}, err
	}
	return input, nil
}

// function to update a task
func (t *TaskService) UpdateTask(number int, input model.Task) (model.Task, error) {
	tasks, err := t.accessTask()
	if err != nil {
		return model.Task{}, err
	}
	//validation input status
	InputStatus := strings.ToLower(input.Status)
	if InputStatus != "" {
		isValid := false
		if InputStatus == utils.NoProgressStatus ||
			InputStatus == utils.OnProgressStatus ||
			InputStatus == utils.CompletedStatus {
			isValid = true
		}
		for i, ts := range tasks {
			if i+1 == number {
				input.ID = ts.ID
				input.Title = ts.Title
				input.Priority = ts.Priority
				tasks[i] = input
				break
			}
		}

		if !isValid {
			return model.Task{}, errors.New("Status is not valid")
		}
	}

	// //validation input priority
	// InputPriority := strings.ToLower(input.Priority)
	// if input.Priority != "" {
	// 	isValid := false
	// 	if InputPriority == utils.LowPriority ||
	// 		InputPriority == utils.MediumPriority ||
	// 		InputPriority == utils.HighPriority {
	// 		isValid = true
	// 	}
	// 	if !isValid {
	// 		return model.Task{}, errors.New("Priority is not valid")
	// 	}
	// }
	// foundtask := false
	// for i := range tasks {
	// 	if tasks[i].ID == input.ID {
	// 		tasks[i].Priority = InputPriority
	// 		foundtask = true
	// 		break
	// 	}
	// }
	// if !foundtask {
	// 	return model.Task{}, utils.ErrNotFOund
	// }
	if err := t.saveTask(tasks); err != nil {
		return model.Task{}, err
	}
	return input, nil
}

// function to delete a task
func (t *TaskService) DeleteTask(number int) error {
	tasks, err := t.accessTask()
	if err != nil {
		return err
	}
	//to initialize new []model.Task chunks efficiently
	newTaskList := make([]model.Task, 0, len(tasks))
	found := false
	for i, ts := range tasks {
		if i+1 == number {
			found = true //penghapusan
			continue
		}
		newTaskList = append(newTaskList, ts)
	}
	if !found {
		return utils.ErrNotFOund
	}
	return t.saveTask(newTaskList)
}

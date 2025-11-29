package utils

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"todo-list-app/model"
)

const FilePath = "data/todos.json"

// check filepath is exist or not and the function will return  error
func IsFIleExist() error {
	_, err := os.Stat(FilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(filepath.Dir(FilePath), 0755); err != nil {
			return err
		}
		return os.WriteFile(FilePath, []byte("[]"), 0644)
	}
	return nil
}

// function will access to read the filepath
func ReadFileJson() ([]model.Task, error) {
	if err := IsFIleExist(); err != nil {
		return nil, err
	}
	bytes, err := os.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}
	var task []model.Task
	if err := json.Unmarshal(bytes, &task); err != nil {
		return nil, errors.New("failed to parse task data: " + err.Error())
	}
	return task, nil
}

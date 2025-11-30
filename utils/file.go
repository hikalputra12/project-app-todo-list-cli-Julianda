package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"todo-list-app/model"
)

const FilePath = "data/todos.json"

// check filepath is exist or not and the function will return  error
func IsFileExist() error {
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
	if err := IsFileExist(); err != nil {
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

// write file
func WriteFileJson(tasks []model.Task) error {
	bytes, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(FilePath, bytes, 0644)
}

// task status
const (
	NoProgressStatus = "noprogress"
	OnProgressStatus = "onprogress"
	CompletedStatus  = "completed"
)

// task priority
const (
	LowPriority    = "low"
	MediumPriority = "medium"
	HighPriority   = "high"
)

// Inisialisasi bufio.Reader SEKALI di scope package
var reader = bufio.NewReader(os.Stdin)

// ReadLine adalah fungsi utilitas yang aman untuk membaca input baris penuh
func ReadLine(prompt string) string {
	fmt.Print(prompt)

	// Membaca seluruh baris hingga karakter newline (\n)
	input, _ := reader.ReadString('\n')

	// Membersihkan spasi di awal/akhir dan karakter newline
	return strings.TrimSpace(input)
}

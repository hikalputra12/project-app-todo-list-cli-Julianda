package cmd

import (
	"fmt"
	"os"
	"strconv"
	"todo-list-app/utils"

	"github.com/spf13/cobra"
)

// Global variable to hold the TaskHandler instance
var taskHandlerInstance *TaskHandler

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A command line interface for managing tasks.",
	Long:  `todo-cli allows you to manage tasks via commands: list, create, update, delete, find.`,
}

// Execute runs the root command. Called from main.go
func Execute(handler *TaskHandler) {
	taskHandlerInstance = handler
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks (ListAllTask).",
	Run: func(cmd *cobra.Command, args []string) {
		// Calls ListAllTask which returns (string, error)
		out, err := taskHandlerInstance.ListAllTask()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(out)
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Interactively creates a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		// Uses utils.ReadLine for safe input
		title := utils.ReadLine("Enter task title: ")
		status := utils.ReadLine("Enter task status (noprogress/onprogress/completed): ")
		priority := utils.ReadLine("Enter task priority (low/medium/high): ")

		// Calls CreateTask which returns (string, error)
		out, err := taskHandlerInstance.CreateTask(title, status, priority)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(out)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates task status by number.",
	Run: func(cmd *cobra.Command, args []string) {
		// Simple logic to get number and status
		numberStr := utils.ReadLine("Enter the number you want to update: ")
		status := utils.ReadLine("status update (noprogress/onprogress/completed): ")

		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("Error: Invalid task number.")
			return
		}

		// Calls UpdateTask which returns (string, error)
		upt, err := taskHandlerInstance.UpdateTask(number, status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(upt)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a task by number.",
	Run: func(cmd *cobra.Command, args []string) {
		// Simple logic to get number
		numberStr := utils.ReadLine("Enter the number of the task you want to delete: ")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println("Error: Invalid task number.")
			return
		}

		// Calls DeleteTask which returns error (message printed internally)
		dltErr := taskHandlerInstance.DeleteTask(number)
		if dltErr != nil {
			// If dltErr is not nil, print the error.
			fmt.Println("Error:", dltErr)
		}
		// Note: If dltErr is nil, DeleteTaskHandler prints success itself
	},
}

var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Finds and displays a task by title.",
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.ReadLine("Enter the title you want to search for: ")

		// Calls GetTaskByTitle which prints the output internally and returns error
		fnErr := taskHandlerInstance.GetTaskByTitle(title)
		if fnErr != nil {
			fmt.Println("Error:", fnErr)
		}
	},
}

// init() adds all subcommands to the rootCmd
func init() {
	rootCmd.AddCommand(listCmd, createCmd, updateCmd, deleteCmd, findCmd)
}

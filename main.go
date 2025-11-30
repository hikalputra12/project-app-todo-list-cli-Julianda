package main

import (
	"fmt"
	"todo-list-app/cmd"
	"todo-list-app/service"
)

func main() {
	// Initialize the service
	taskService := service.NewTaskService()
	// Initialize the handler with the service
	taskHandler := cmd.NewTaskHandler(taskService)

	for {
		fmt.Println("=== Task CLI ===")
		fmt.Println("1. List tasks")
		fmt.Println("2. Create task")
		fmt.Println("3. Update task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Get task by title")
		fmt.Println("0. Exit")
		fmt.Print("Choose menu: ")

		var choice string
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			out, err := taskHandler.ListAllTask()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)
		case "2":
			var title, status, priority string
			fmt.Print("Enter task title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter task status (noprogress/onprogress/completed): ")
			fmt.Scanln(&status)
			fmt.Print("Enter task priority (low/medium/high): ")
			fmt.Scanln(&priority)
			out, err := taskHandler.CreateTask(title, status, priority)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)
		case "3":
			out, err := taskHandler.ListAllTask()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)
			var id int
			var status string
			fmt.Print("masukkan nomor yang ingin anda update: ")
			fmt.Scanln(&id)
			fmt.Print("update status (no progress/on progress/completed): ")
			fmt.Scanln(&status)
			upt, err := taskHandler.UpdateTask(id, status)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(upt)
		case "4":
			out, err := taskHandler.ListAllTask()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println(out)
			var number int
			fmt.Print("masukkan nomor tugas yang anda ingin hapus: ")
			fmt.Scanln(&number)
			dlt := taskHandler.DeleteTask(number)
			fmt.Println(dlt)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		fmt.Println() // Print a new line for better readability
	}
}

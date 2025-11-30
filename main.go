package main

import (
	"fmt"
	"todo-list-app/cmd"
	"todo-list-app/service"
	"todo-list-app/utils"
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
			title := utils.ReadLine("Enter task title: ")
			fmt.Scanln(&title)
			status := utils.ReadLine("Enter task status (noprogress/onprogress/completed): ")
			fmt.Scanln(&status)
			priority := utils.ReadLine("Enter task priority (low/medium/high): ")
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
			var num int
			var status string
			fmt.Print("Enter the number you want to update: ")
			fmt.Scanln(&num)
			fmt.Print("status update (noprogress/onprogress/completed):")
			fmt.Scanln(&status)
			upt, err := taskHandler.UpdateTask(num, status)
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
		case "5":
			title := utils.ReadLine("masukkan judul yang anda ingin cari: ")
			fmt.Scanln(&title)
			fn := taskHandler.GetTaskByTitle(title)
			fmt.Println(fn)

		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
		fmt.Println() // Print a new line for better readability
	}
}

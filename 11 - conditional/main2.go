package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
	DueDate   time.Time
}

func main() {
	var tasks []Task

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1 - Add task")
		fmt.Println("2 - List tasks")
		fmt.Println("3 - Mark task as completed")
		fmt.Println("0 - Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title string
			var dueDateStr string

			fmt.Print("Enter task title: ")
			fmt.Scan(&title)
			fmt.Print("Enter due date (YYYY-MM-DD): ")
			fmt.Scan(&dueDateStr)

			dueDate, _ := time.Parse("2006-01-02", dueDateStr)
			task := Task{ID: len(tasks) + 1, Title: title, DueDate: dueDate}
			tasks = append(tasks, task)
			fmt.Println("Task added successfully!")

		case 2:
			fmt.Println("Tasks:")
			for _, task := range tasks {
				completedStatus := "No"
				if task.Completed {
					completedStatus = "Yes"
				}
				fmt.Printf("%d - %s (Completed: %s)\n", task.ID, task.Title, completedStatus)
			}

		case 3:
			fmt.Print("Enter task ID: ")
			var taskID int
			fmt.Scan(&taskID)

			found := false
			for i := range tasks {
				if tasks[i].ID == taskID {
					tasks[i].Completed = true
					fmt.Println("Task marked as completed!")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Task not found!")
			}
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

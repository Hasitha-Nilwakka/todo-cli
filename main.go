package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	//check if the user has provided sufficient arguments
	if len(args) < 2 {
		fmt.Println("Usage: ./todo [add|list|complete|delete] [arguments]")
		return
	}

	//main functional argument of the user
	cmd := args[1]
	switch cmd {
	case "add": //adding a new item
		if len(args) < 3 {
			fmt.Println("Usage: ./todo add \"task description\"")
			return
		}
		title := args[2]
		if err := AddTask(title); err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println("Task added successfully!")
	case "list": //display a list of records
		err := ListTasks()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	case "complete": //mark an item as complete
		if len(args) < 3 {
			fmt.Println("Usage: ./todo add \"task number\"")
			return
		}
		taskId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err := CompleteTask(taskId); err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println("Task marked as completed!")
	case "delete": //delete a task
		if len(args) < 3 {
			fmt.Println("Usage: ./todo delete \"task number\"")
			return
		}
		taskId, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid task id")
		}

		if err := DeleteTask(taskId); err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println("Task deleted successfully!")
	default:
		//if user argument is incorrect, usage is displayed
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Usage: todo [add|list|complete|delete]")
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// The main data structre for the application
type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// function loads the current tasks in the JSON file.
// if JSON is not found a nil strcut is returned
func Loadtasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{}, nil
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}

// addtask function prepares data to be saved to the JSON
// generate the task id by incrementing the last task id
func AddTask(title string) error {
	tasks, err := Loadtasks()
	if err != nil {
		return err
	}
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	tasks = append(tasks, Task{ID: id, Title: title, Completed: false})
	return Savetask(tasks)
}

// function takes in a []Task and save the data into JSON
func Savetask(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

// the function markes a user defined tasks as completed.
// the function calls SaveTask function to update the tasks to JSON
func CompleteTask(taskId int) error {
	tasks, err := Loadtasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == taskId {
			tasks[i].Completed = true
			return Savetask(tasks)
		}
	}

	return fmt.Errorf("task ID %d not found", taskId)
}

// calls load tasks to retrive the current list of tasks
// Status is marked as compled has "✅" and incomplete has "❌"
// formatted records are printed to console line by line
func ListTasks() error {
	tasks, err := Loadtasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d %s %s\n", task.ID, task.Title, status)
	}

	return nil
}

// This function deletes a user define task id from the JSON records
// Other tasks are then recorded to a new struct and calls SaveTask function
func DeleteTask(taskId int) error {
	tasks, err := Loadtasks()
	if err != nil {
		return nil
	}
	newTasks := []Task{}
	found := false
	for _, task := range tasks {
		if task.ID == taskId {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		return fmt.Errorf("task id %d not found", taskId)
	}

	return Savetask(newTasks)
}

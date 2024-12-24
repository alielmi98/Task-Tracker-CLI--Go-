package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alielmi98/Task-Tracker-CLI-Go-Implementation/internal/app"
)

func main() {
	filePath := filepath.Join("..", "..", "data", "tasks.json")
	taskManager, err := app.NewTaskManager(filePath)
	if err != nil {
		fmt.Printf("Error creating task manager: %v\n", err)
		return
	}

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	markInProgressCmd := flag.NewFlagSet("mark-in-progress", flag.ExitOnError)
	markDoneCmd := flag.NewFlagSet("mark-done", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Please specify a command (add, update, delete, list, mark-in-progress, mark-done)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		description := strings.Join(addCmd.Args(), " ")
		taskManager.AddTask(description)
	case "update":
		updateCmd.Parse(os.Args[2:])
		args := updateCmd.Args()
		if len(args) < 2 {
			fmt.Println("Usage: update <id> <new-description>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		newDesc := strings.Join(args[1:], " ")
		taskManager.UpdateTask(id, newDesc)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		args := deleteCmd.Args()
		if len(args) < 1 {
			fmt.Println("Usage: delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		taskManager.DeleteTask(id)
	case "list":
		listCmd.Parse(os.Args[2:])
		args := listCmd.Args()
		if len(args) > 0 {
			if isValidStatus(args[0]) {
				taskManager.ListFilterByStatus(args[0])
			} else {
				fmt.Printf("invalid filtering input. Valid inputs are: todo, done, in-progress\nUsage: list [status]\n")
				os.Exit(1)
			}
		} else {
			taskManager.ListTasks()
		}

	case "mark-in-progress":
		markInProgressCmd.Parse(os.Args[2:])
		args := markInProgressCmd.Args()
		if len(args) < 1 {
			fmt.Println("Usage: mark-in-progress <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		taskManager.MarkTaskInProgress(id)
	case "mark-done":
		markDoneCmd.Parse(os.Args[2:])
		args := markDoneCmd.Args()
		if len(args) < 1 {
			fmt.Println("Usage: mark-done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		taskManager.MarkTaskDone(id)
	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

func isValidStatus(status string) bool {
	validStatuses := []string{"todo", "done", "in-progress"}
	for _, v := range validStatuses {
		if v == status {
			return true
		}
	}
	return false
}

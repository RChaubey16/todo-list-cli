package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RChaubey16/todo-list-cli/internal/tasks"
)

func main() {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "task description")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeID := removeCmd.Int("id", 0, "task ID to remove")

	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	updateID := updateCmd.Int("id", 0, "task ID to update")

	// No subcommand?
	if len(os.Args) < 2 {
		fmt.Println("expected 'add', 'list', or 'remove'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTask == "" {
			fmt.Println("missing --task")
			os.Exit(1)
		}
		err := tasks.AddTaskToList(*addTask)
		if err != nil {
			fmt.Println("Error occurred while adding the task", err)
		}
		fmt.Println("Task added to the list:", *addTask)

	case "list":
		listCmd.Parse(os.Args[2:])
		tasks.ListTasks()

	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *removeID == 0 {
			fmt.Println("missing --id")
			os.Exit(1)
		}
		tasks.RemoveTaskFromList(*removeID)
		fmt.Println("Task removed from list:", *removeID)

	case "update":
		updateCmd.Parse(os.Args[2:])
		if *updateID == 0 {
			fmt.Println("missing --id")
			os.Exit(1)
		}
		err := tasks.UpdateTask(*updateID)
		if err != nil {
			fmt.Println("Failed to update task:", err)
		}
		fmt.Println("Task updated:", *updateID)

	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}
}

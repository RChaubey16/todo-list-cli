package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "task description")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	removeID := removeCmd.Int("id", 0, "task ID to remove")

	fmt.Println("OS ARGS", len(os.Args), os.Args)

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
		fmt.Println("Adding task:", *addTask)

	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Listing tasks...")

	case "remove":
		removeCmd.Parse(os.Args[2:])
		if *removeID == 0 {
			fmt.Println("missing --id")
			os.Exit(1)
		}
		fmt.Println("Removing task with ID:", *removeID)

	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}
}

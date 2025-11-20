package tasks

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RChaubey16/todo-list-cli/internal/types"
)

func LoadTasks(fileName string) (*types.TaskFile, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read the tasks file: %w", err)
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return nil, fmt.Errorf("invalid json: %w", err)
	}

	return &tf, nil
}

func ListTasks() error {
	taskFile, err := LoadTasks("output.json")
	if err != nil {
		return err
	}

	if len(taskFile.Tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}

	for i, t := range taskFile.Tasks {
		fmt.Printf("%d. [%s] %s\n", i+1, t.Status, t.Title)
	}

	return nil
}

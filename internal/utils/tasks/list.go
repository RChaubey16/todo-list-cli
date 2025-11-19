package tasks

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskItem struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TaskFile struct {
	Tasks []TaskItem `json:"tasks"`
}

func LoadTasks(fileName string) (*TaskFile, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read the tasks file: %w", err)
	}

	var tf TaskFile
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

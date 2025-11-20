package tasks

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RChaubey16/todo-list-cli/internal/types"
)

func AddTaskToList(task string) error {
	newTask := types.TaskItem{
		Title:  task,
		Status: "pending",
	}

	data, err := os.ReadFile("output.json")
	if err != nil {
		return fmt.Errorf("could not read the tasks file: %w", err)
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	tf.Tasks = append(tf.Tasks, newTask)

	updatedJSON, err := json.MarshalIndent(tf, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("output.json", updatedJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

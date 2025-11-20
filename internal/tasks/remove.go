package tasks

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RChaubey16/todo-list-cli/internal/types"
)

func RemoveTaskFromList(taskId int) error {
	data, err := os.ReadFile("output.json")
	if err != nil {
		return fmt.Errorf("could not read the tasks file: %w", err)
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	indexToRemove := taskId - 1
	if indexToRemove <= 0 {
		return fmt.Errorf("Invalid task ID:", taskId)
	}

	tf.Tasks = append(tf.Tasks[:indexToRemove], tf.Tasks[indexToRemove+1:]...)

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

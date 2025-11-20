package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/RChaubey16/todo-list-cli/internal/types"
	"github.com/RChaubey16/todo-list-cli/internal/utils"
)

func RemoveTaskFromList(taskId int) error {
	data, err := utils.ReadJsonFile("output.json")
	if err != nil {
		return err
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	indexToRemove := taskId - 1

	if indexToRemove < 0 || indexToRemove >= len(tf.Tasks) {
		return fmt.Errorf("invalid task ID")
	}

	tf.Tasks = append(tf.Tasks[:indexToRemove], tf.Tasks[indexToRemove+1:]...)

	err = utils.WriteToJsonFile(tf, "output.json")
	if err != nil {
		return err
	}

	return nil
}

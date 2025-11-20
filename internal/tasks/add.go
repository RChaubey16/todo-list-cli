package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/RChaubey16/todo-list-cli/internal/types"
	"github.com/RChaubey16/todo-list-cli/internal/utils"
)

func AddTaskToList(task string) error {
	data, err := utils.ReadJsonFile("output.json")
	if err != nil {
		return err
	}

	newTask := types.TaskItem{
		Title:  task,
		Status: "pending",
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	tf.Tasks = append(tf.Tasks, newTask)

	err = utils.WriteToJsonFile(tf, "output.json")
	if err != nil {
		return err
	}

	return nil
}

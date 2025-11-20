package tasks

import (
	"encoding/json"
	"fmt"

	"github.com/RChaubey16/todo-list-cli/internal/types"
	"github.com/RChaubey16/todo-list-cli/internal/utils"
)

func UpdateTask(taskId int) error {
	data, err := utils.ReadJsonFile("output.json")
	if err != nil {
		return err
	}

	var tf types.TaskFile
	if err := json.Unmarshal(data, &tf); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	idx := taskId - 1

	if idx < 0 || idx >= len(tf.Tasks) {
		fmt.Println("invalid task index")
		return nil
	}

	tf.Tasks[idx].Status = "done"

	err = utils.WriteToJsonFile(tf, "output.json")
	if err != nil {
		return err
	}

	return nil
}

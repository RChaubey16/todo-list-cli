package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/RChaubey16/todo-list-cli/internal/types"
)

func ReadJsonFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("could not read the tasks file: %w", err)
	}

	return data, nil
}

func WriteToJsonFile(fileContent types.TaskFile, fileName string) error {
	updatedJSON, err := json.MarshalIndent(fileContent, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, updatedJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

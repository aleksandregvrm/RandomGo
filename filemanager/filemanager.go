package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error with opening relevant file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to read line in file")
	}
	return lines, nil
}

func WriteJSON(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create a file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to JSON")
	}
	return nil
}

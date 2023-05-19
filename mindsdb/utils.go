package mindsdb

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LoadEnvFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Error(message string) error {
	return errors.New(message)
}

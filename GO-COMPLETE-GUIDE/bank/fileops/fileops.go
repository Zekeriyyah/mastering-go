package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(filename string, value float64) error {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(filename, []byte(valueText), 0644)
	// err := errors.New("dummy error")
	return err
}

func GetFloatFromFile(filename string) (float64, error) {
	valueByte, err := os.ReadFile(filename)
	if err != nil {
		return 0, errors.New("user record not found")
	}

	value, err := strconv.ParseFloat(string(valueByte), 64)
	if err != nil {
		return 0, errors.New("failed to parse value")
	}

	return value, nil
}

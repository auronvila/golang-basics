package fileOps

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// set the first letter to uppercase to export the function to the project that you will use
func WriteFloatToFile(val float64, filename string) {
	valueText := fmt.Sprint(val)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("failed to find the file")
	}
	valueText := file
	balance, err := strconv.ParseFloat(string(valueText), 64)
	if err != nil {
		return 1000, errors.New("failed to parse the stored val")
	}
	return balance, nil
}

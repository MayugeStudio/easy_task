package code

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(fileName string) (lines []string, err error) {
	file, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("Error opening file:", openErr)
		os.Exit(1)
	}
	defer func(f *os.File) {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanErr := scanner.Err(); scanErr != nil {
		return nil, scanErr
	}
	return lines, nil
}

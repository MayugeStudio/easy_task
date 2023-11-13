package code

import (
	"bufio"
	"fmt"
	"os"
)

func ScanFile(fileName string) []string {
	file, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("Error opening file:", openErr)
		os.Exit(1)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err.Error())
			os.Exit(1)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		os.Exit(1)
	}
	return lines
}

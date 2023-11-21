package code

import (
	"bufio"
	"os"
)

type FileReader interface {
	ReadLines(filename string) (lines []string, err error)
}

type FileScanner struct{}

func (fs FileScanner) ReadLines(filename string) (lines []string, err error) {
	file, openErr := os.Open(filename)
	if openErr != nil {
		return nil, openErr
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

func ScanFile(fileName string, reader FileReader) (lines []string, err error) {
	return reader.ReadLines(fileName)
}

package src

import (
	"bufio"
	"os"
)

type FileReader interface {
	ReadLines(fileName string) (lines []string, err error)
}

type FileScanner struct{}

func (fs FileScanner) ReadLines(fileName string) (lines []string, err error) {
	file, openErr := os.Open(fileName)
	if openErr != nil {
		return nil, openErr
	}
	defer func(f *os.File) {
		_ = f.Close()
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

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type TokenType string

const (
	PriorityHigh   TokenType = "H"
	PriorityMedium           = "M"
	PriorityLow              = "L"
)

var (
	ContinueLine  = errors.New("skip line")
	InvalidSyntax = errors.New("invalid file structure")
)
var logger = log.New(os.Stderr, "", log.Lmsgprefix)

type Task struct {
	Title    string
	IsDone   bool
	Priority string
}

func NewTask() *Task {
	return &Task{
		Title:    "",
		IsDone:   false,
		Priority: "",
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		logger.Fatal("Usage: tct [filename]")
	}
	fileName := flag.Arg(0)
	lines := scanFile(fileName)
	tasks := parseLines(lines)
	printTasks(tasks)
}

func scanFile(fileName string) []string {
	file, openErr := os.Open(fileName)
	if openErr != nil {
		fmt.Println("Error opening file:", openErr)
		os.Exit(1)
	}
	defer closeFile(file)

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

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		logger.Fatal("Error closing file: ", err)
	}
}

func parseLines(lines []string) []*Task {
	tasks := make([]*Task, 0)
	for _, line := range lines {
		if task, err := parseLine(line); err != nil {
			if errors.Is(err, ContinueLine) {
				continue
			} else {
				fmt.Printf("Error parsing task: %s\n", err.Error())
				fmt.Printf("\tLine: %s\n", line)
				continue
			}
		} else {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func parseLine(line string) (*Task, error) {
	task := NewTask()
	if !strings.HasPrefix(line, "-") {
		return nil, ContinueLine
	}
	line = strings.TrimPrefix(line, "-")
	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")
	tokens := strings.Fields(line)
	if len(tokens) == 0 {
		return nil, InvalidSyntax
	} else if len(tokens) <= 3 {
		return nil, InvalidSyntax
	}
parsing:
	for {
		token := strings.ToUpper(tokens[0])
		switch TokenType(token) {
		case "X":
			task.IsDone = true
		case PriorityHigh:
			task.Priority = "High"
		case PriorityMedium:
			task.Priority = "Medium"
		case PriorityLow:
			task.Priority = "Low"
		default:
			task.Title = strings.Join(tokens, " ")
			break parsing
		}
		tokens = tokens[1:]
	}
	return task, nil
}

func printTasks(tasks []*Task) {
	for i, task := range tasks {
		printTask(i, task)
	}
}

func printTask(index int, task *Task) {
	var doneStr string
	if task.IsDone {
		doneStr = "Complete"
	} else {
		doneStr = "InProgress"
	}
	fmt.Printf("[%d] %s\n", index, task.Title)
	fmt.Printf("\tPriority: %s\n", task.Priority)
	fmt.Printf("\tStatus: %s\n", doneStr)
}

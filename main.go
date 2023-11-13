package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	High   = "H"
	Medium = "M"
	Low    = "L"
)

var PriorityMap = map[string]string{
	High:   "High",
	Medium: "Medium",
	Low:    "Low",
}

var (
	ContinueLine      = errors.New("skip line")
	InvalidLineSyntax = errors.New("invalid file structure")
)

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
	args := os.Args[1:]
	if len(args) >= 2 {
		log.Fatal("Usage: tct [filename]")
	} else if len(args) == 0 {
		fmt.Println("Error: need one argument.")
		fmt.Println("Usage: tct [filename]")
		os.Exit(1)
	}
	fileName := args[0]
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
		fmt.Println("Error closing file:", err)
		os.Exit(1)
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
	fields := strings.Fields(line)

	if len(fields) <= 3 {
		return nil, InvalidLineSyntax
	}
parsing:
	for {
		field := strings.ToUpper(fields[0])
		switch field {
		case "X":
			task.IsDone = true
		case High:
			task.Priority = PriorityMap[High]
		case Medium:
			task.Priority = PriorityMap[Medium]
		case Low:
			task.Priority = PriorityMap[Low]
		default:
			task.Title = strings.Join(fields, " ")
			break parsing
		}
		fields = fields[1:]
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

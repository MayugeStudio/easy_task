package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type TokenType string

const (
	PriorityHigh   TokenType = "H"
	PriorityMedium           = "M"
	PriorityLow              = "L"
)

const (
	DoneSymbol     = "X"
	UndoneSymbol   = " "
	ProgressSymbol = "#"
)

var InvalidSyntax = errors.New("invalid file structure")

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
		fmt.Printf("Usage: tst [filename]")
		os.Exit(1)
	}
	fileName := flag.Arg(0)
	lines := scanFile(fileName)
	tasks, msgSlice := parseLines(lines)
	printErrorMessages(msgSlice)
	printTasks(tasks)
	printTaskProgress(tasks)
}

func scanFile(fileName string) []string {
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

func parseLines(lines []string) ([]*Task, []string) {
	tasks := make([]*Task, 0)
	errMsgSlice := make([]string, 0)
	for i, line := range lines {
		pLine, skip := toProcessedLineFromRawLine(line)
		if skip {
			continue
		}
		tokens, err := processedLineToTokens(pLine)
		if err != nil {
			msg := fmt.Sprintf("Error in preprocessing task: %s\n", err.Error())
			msg += fmt.Sprintf("  > in line - %d\n", i+1)
			msg += fmt.Sprintf("     > %q", line)
			errMsgSlice = append(errMsgSlice, msg)
			continue
		}
		task := parseLine(tokens)
		tasks = append(tasks, task)
	}
	return tasks, errMsgSlice
}

func parseLine(tokens []string) *Task {
	task := NewTask()
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
	return task
}

func toProcessedLineFromRawLine(line string) (processedLine string, skip bool) {
	if !strings.HasPrefix(line, "-") {
		return "", true
	}
	line = strings.TrimPrefix(line, "-")
	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")
	return line, false
}

func processedLineToTokens(processedLine string) ([]string, error) {
	tokens := strings.Fields(processedLine)
	if len(tokens) == 0 {
		return nil, InvalidSyntax
	} else if len(tokens) <= 3 {
		return nil, InvalidSyntax
	}
	return tokens, nil
}

func printErrorMessages(messages []string) {
	if len(messages) == 0 {
		return
	}
	for _, message := range messages {
		fmt.Print(message)
	}
	fmt.Printf("\n")
}

func getMaxTaskNameLength(tasks []*Task) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

func printTasks(tasks []*Task) {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		printTask(task, maxTaskNameLength)
	}
}

func printTask(task *Task, maxTaskNameLength int) {
	var doneStr string
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	fmt.Printf("[%s] %-*s ~ <priority: %s>\n", doneStr, maxTaskNameLength, task.Title, task.Priority)
}

func printTaskProgress(tasks []*Task) {
	progressBarLength := 20.0
	taskNum := float64(len(tasks))
	doneTaskNum := 0.0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	doneTaskRatio := doneTaskNum / taskNum
	doneTaskStrLength := int(doneTaskRatio * progressBarLength)
	doneTaskStr := strings.Repeat(ProgressSymbol, doneTaskStrLength)
	undoneTaskStr := strings.Repeat(" ", int(progressBarLength)-doneTaskStrLength)
	fmt.Printf("[%s%s]%d%%", doneTaskStr, undoneTaskStr, int(doneTaskRatio*100))
}

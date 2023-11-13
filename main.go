package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"math"
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
		fmt.Printf("Error closing file: %s\n", err.Error())
		os.Exit(1)
	}
}

func parseLines(lines []string) ([]*Task, []string) {
	tasks := make([]*Task, 0)
	errMsgSlice := make([]string, 0)
	for i, line := range lines {
		pLine, skip := toPreprocessedLine(line)
		if skip {
			continue
		}
		tokens, err := preprocessedLineToTokens(pLine)
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

func toPreprocessedLine(line string) (preprocessedLine string, skip bool) {
	if !strings.HasPrefix(line, "-") {
		return "", true
	}
	line = strings.TrimPrefix(line, "-")
	line = strings.ReplaceAll(line, "[", "")
	line = strings.ReplaceAll(line, "]", "")
	return line, false
}

func preprocessedLineToTokens(line string) ([]string, error) {
	tokens := strings.Fields(line)
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

func printTasks(tasks []*Task) {
	maxTaskNameLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxTaskNameLength {
			maxTaskNameLength = len(task.Title)
		}
	}
	for _, task := range tasks {
		printTask(task, maxTaskNameLength)
	}
}

func printTask(task *Task, maxTaskNameLength int) {
	var doneStr string
	var paddingStr string
	if len(task.Title) <= maxTaskNameLength {
		paddingStr = strings.Repeat(" ", maxTaskNameLength-len(task.Title))
	}
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	fmt.Printf("[%s] %s%s ~ <priority: %s>\n", doneStr, task.Title, paddingStr, task.Priority)
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
	doneTaskStrLength := int(math.Ceil(progressBarLength * doneTaskRatio))
	doneTaskStr := strings.Repeat(ProgressSymbol, doneTaskStrLength)
	undoneTaskStr := strings.Repeat(" ", int(progressBarLength)-doneTaskStrLength)
	fmt.Printf("[%s%s]", doneTaskStr, undoneTaskStr)
}

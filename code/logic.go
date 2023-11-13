package code

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type TokenType string

const (
	DoneSymbol     = "X"
	UndoneSymbol   = " "
	ProgressSymbol = "#"
)

const DefaultProgressBarLength = 40.0

var InvalidSyntax = errors.New("invalid file structure")

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

func ParseLines(lines []string) ([]TaskPtr, []string) {
	tasks := make([]TaskPtr, 0)
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

func parseLine(tokens []string) TaskPtr {
	task := NewTask()
parsing:
	for {
		token := strings.ToUpper(tokens[0])
		switch TokenType(token) {
		case "X":
			task.IsDone = true
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
		return nil, InvalidSyntax // 'X TaskTitle' or '  TaskTitle'
	}
	return tokens, nil
}

func PrintErrorMessages(messages []string) {
	if len(messages) == 0 {
		return
	}
	for _, message := range messages {
		fmt.Print(message)
	}
	fmt.Printf("\n")
}

func getMaxTaskNameLength(tasks []TaskPtr) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

func PrintTasks(tasks []TaskPtr) {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		printTask(task, maxTaskNameLength)
	}
}

func printTask(task TaskPtr, maxTaskNameLength int) {
	var doneStr string
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	fmt.Printf("[%s] %-*s\n", doneStr, maxTaskNameLength, task.Title)
}

func PrintTaskProgress(tasks []TaskPtr) {
	progressBarLength := DefaultProgressBarLength
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

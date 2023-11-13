package code

import (
	"fmt"
	"strings"
)

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
	for {
		isDone := processLine(tokens, task)
		if isDone {
			break
		}
		tokens = tokens[1:]
	}
	return task
}

func processLine(tokens []string, task TaskPtr) (done bool) {
	token := strings.ToUpper(tokens[0])
	switch token {
	case "X":
		task.IsDone = true
		return false
	default:
		task.Title = strings.Join(tokens, " ")
		return true
	}
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

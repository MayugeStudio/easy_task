package code

import (
	"fmt"
	"strings"
)

func ParseLines(lines []string) ([]TaskPtr, []string) {
	tasks := make([]TaskPtr, 0)
	errorMessages := make([]string, 0)
	for i, line := range lines {
		pLine, skip := removeUnnecessaryTokenFromLine(line)
		if skip {
			continue
		}
		tokens, err := processedLineToTokens(pLine)
		if err != nil {
			msg := fmt.Sprintf("Error in preprocessing task: %s\n", err.Error())
			msg += fmt.Sprintf("  > in line - %d\n", i+1)
			msg += fmt.Sprintf("     > %q", line)
			errorMessages = append(errorMessages, msg)
			continue
		}
		task := parseLine(tokens)
		tasks = append(tasks, task)
	}
	return tasks, errorMessages
}

func parseLine(tokens []string) TaskPtr {
	task := NewTask()
parsing:
	for {
		token := strings.ToUpper(tokens[0])
		switch token {
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

func removeUnnecessaryTokenFromLine(line string) (processedLine string, skip bool) {
	if !strings.HasPrefix(line, "-") {
		return "", true
	}
	line = strings.TrimPrefix(line, "-")
	line = strings.Replace(line, "[", "", 1)
	line = strings.Replace(line, "]", "", 1)
	return line, false
}

func processedLineToTokens(processedLine string) ([]string, error) {
	tokens := strings.Fields(processedLine)
	if len(tokens) == 0 {
		return nil, InvalidSyntax // 'X TaskTitle' or '  TaskTitle'
	}
	return tokens, nil
}

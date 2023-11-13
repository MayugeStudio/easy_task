package code

import (
	"strings"
)

func ParseLines(lines []string) []TaskPtr {
	tasks := make([]TaskPtr, 0)
	for _, line := range lines {
		pLine, skip := removeUnnecessaryTokenFromLine(line)
		if skip {
			continue
		}
		tokens, skip := processedLineToTokens(pLine)
		if skip {
			continue
		}
		task := parseLine(tokens)
		tasks = append(tasks, task)
	}
	return tasks
}

func parseLine(tokens []string) TaskPtr {
	task := NewTask()
	for tokens != nil {
		token := strings.ToUpper(tokens[0])
		switch token {
		case "X":
			task.IsDone = true
		default:
			task.Title = strings.Join(tokens, " ")
			tokens = nil
		}
		if tokens != nil {
			tokens = tokens[1:]
		}
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

func processedLineToTokens(processedLine string) (tokens []string, skip bool) {
	tokens = strings.Fields(processedLine)
	if len(tokens) == 0 {
		return nil, true
	}
	return tokens, false
}

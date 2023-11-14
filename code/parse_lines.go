package code

import (
	"strings"
)

// ParseLines processes a slice of strings representing tasks
// and returns a slice of Task pointers
func ParseLines(lines []string) []TaskPtr {
	tasks := make([]TaskPtr, 0)
	for _, line := range lines {
		if !isLineValid(line) {
			continue
		}
		cleanLine := processLine(line)
		tokens := strings.Fields(cleanLine)
		task := convertTokensToTask(tokens)
		tasks = append(tasks, task)
	}
	return tasks
}

// convertTokensToTask takes a slice of strings representing tokens
// and convert them into a Task pointer.
func convertTokensToTask(tokens []string) TaskPtr {
	task := NewTask()
	// Process each token until the tokens slice is empty.
	for len(tokens) > 0 {
		token := strings.ToUpper(tokens[0])
		switch token {
		case "X":
			task.IsDone = true
		default:
			task.Title = strings.Join(tokens, " ")
			tokens = nil
		}
		if len(tokens) > 0 {
			tokens = tokens[1:]
		}
	}
	return task
}

// isLineValid checks if a line meets the criteria for a valid task.
func isLineValid(line string) bool {
	noBlankCharCount := len(line) - strings.Count(line, " ")
	return len(line) >= 2 && noBlankCharCount >= 2 && strings.HasPrefix(line, "-")
}

// processLine removes unnecessary tokens from a line.
func processLine(line string) string {
	line = strings.TrimPrefix(line, "-")
	line = strings.Replace(line, "[", "", 1)
	line = strings.Replace(line, "]", "", 1)
	return line
}

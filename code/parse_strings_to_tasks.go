package code

import "strings"

// ParseStringsToTasks processes a slice of strings representing tasks
// and returns a slice of Task pointers
func ParseStringsToTasks(taskStrings []string) []*Task {
	tasks := make([]*Task, 0)
	taskStrings = FormatTaskStrings(taskStrings)
	for _, str := range taskStrings {
		if strings.HasPrefix(str, "-") {
			str = strings.TrimPrefix(str, "-")
			str = strings.TrimSpace(str)
			str = strings.Replace(str, "[", "", 1)
			str = strings.Replace(str, "]", "", 1)
			str = strings.TrimSpace(str)
			tokens := strings.Fields(str)
			task := ParseStringToTask(tokens)
			tasks = append(tasks, task)
		} else {
			continue
		}
	}
	return tasks
}

// ParseStringToTask takes a slice of strings representing tokens
// and convert them into a Task pointer.
func ParseStringToTask(tokens []string) *Task {
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

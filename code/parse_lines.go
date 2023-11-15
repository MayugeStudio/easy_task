package code

import "strings"

// ParseTaskStringsToTasks processes a slice of strings representing tasks
// and returns a slice of Task pointers
func ParseTaskStringsToTasks(taskStrings []string) []*Task {
	tasks := make([]*Task, 0)
	taskStrings = FormatTaskStrings(taskStrings)
	for _, str := range taskStrings {
		if strings.HasPrefix(str, "-") {
			tokens := strings.Fields(str)
			task := ParseTaskStringToTask(tokens)
			tasks = append(tasks, task)
		} else {
			continue
		}
	}
	return tasks
}

// ParseTaskStringToTask takes a slice of strings representing tokens
// and convert them into a Task pointer.
func ParseTaskStringToTask(tokens []string) *Task {
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

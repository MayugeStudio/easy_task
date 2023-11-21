package code

import "strings"

func ParseStringsToTasks(taskStrings []string) *TodoItemContainer {
	todoItemContainer := NewTodoItemContainer()
	taskStrings = FormatTaskStrings(taskStrings)
	for _, str := range taskStrings {
		if IsSingleTaskString(str) {
			task := parseSingleTaskString(str)
			todoItemContainer.AddTask(task)
			continue
		}

		if IsGroupTitle(str) {
			continue
		}

		if IsGroupTaskString(str) {
			continue
		}
	}
	return todoItemContainer
}

func parseSingleTaskString(str string) *Task {
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	str = strings.TrimSpace(str)
	tokens := strings.Fields(str)
	title := ""
	isDone := false
	// Process each token until the tokens slice is empty.
	for len(tokens) > 0 {
		token := tokens[0]
		switch token {
		case "X":
			isDone = true
		default:
			title = strings.Join(tokens, " ")
			tokens = nil
		}
		if len(tokens) > 0 {
			tokens = tokens[1:]
		}
	}
	task := NewTask(title, isDone)
	return task
}

package code

import "strings"

func ParseStringsToTasks(taskStrings []string) *TodoItemContainer {
	todoItemContainer := NewTodoItemContainer()
	taskStrings = FormatTaskStrings(taskStrings)
	var group *Group
	for _, str := range taskStrings {
		if IsSingleTaskString(str) {
			task := parseSingleTaskString(str)
			todoItemContainer.AddTask(task)
			continue
		}

		if IsGroupTitle(str) {
			group = parseGroupTaskTitle(str)
			todoItemContainer.AddGroup(group)
			continue
		}

		if IsGroupTaskString(str) {
			task := parseGroupTaskString(str)
			if group != nil {
				group.AddTask(task)
			}
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

func parseGroupTaskTitle(str string) *Group {
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := NewGroup(str)
	return g
}

func parseGroupTaskString(str string) *Task {
	title := ""
	isDone := false

	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "-")
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	str = strings.TrimSpace(str)

	if strings.HasPrefix(str, "X") || strings.HasPrefix(str, "x") {
		isDone = true
		if strings.HasPrefix(str, "X") {
			str = strings.TrimPrefix(str, "X")
		} else if strings.HasPrefix(str, "x") {
			str = strings.TrimPrefix(str, "x")
		}
		str = strings.TrimSpace(str)
	}

	title = str
	task := NewTask(title, isDone)
	return task
}

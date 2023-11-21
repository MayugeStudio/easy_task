package src

import "strings"

func ParseStringsToTasks(taskStrings []string) *TodoItemContainer {
	todoItemContainer := NewTodoItemContainer()
	taskStrings = FormatTaskStrings(taskStrings)
	var group *Group
	for _, str := range taskStrings {
		if IsSingleTaskString(str) {
			task := parseTaskString(str)
			todoItemContainer.AddTask(task)
			continue
		}

		if IsGroupTitle(str) {
			group = parseGroupTaskTitle(str)
			todoItemContainer.AddGroup(group)
			continue
		}

		if IsGroupTaskString(str) {
			str = strings.TrimSpace(str)
			task := parseTaskString(str)
			if group != nil {
				group.AddTask(task)
			}
			continue
		}
	}
	return todoItemContainer
}

func parseTaskString(str string) *Task {
	title := ""
	isDone := false
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	str = strings.Replace(str, "[", "", 1)
	str = strings.Replace(str, "]", "", 1)
	str = strings.TrimSpace(str)
	if strings.HasPrefix(str, "X") {
		isDone = true
		str = strings.TrimPrefix(str, "X")
		str = strings.TrimSpace(str)
	}

	title = str

	task := NewTask(title, isDone)
	return task
}

func parseGroupTaskTitle(str string) *Group {
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := NewGroup(str)
	return g
}

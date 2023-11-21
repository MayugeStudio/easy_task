package logic

import (
	"easy_task/src/domain"
	"strings"
)

func ParseStringsToTasks(taskStrings []string) *domain.TodoItemContainer {
	todoItemContainer := domain.NewTodoItemContainer()
	taskStrings = FormatTaskStrings(taskStrings)
	var group *domain.Group
	for _, str := range taskStrings {
		if isSingleTaskString(str) {
			task := parseTaskString(str)
			todoItemContainer.AddTask(task)
			continue
		}

		if isGroupTitle(str) {
			group = parseGroupTaskTitle(str)
			todoItemContainer.AddGroup(group)
			continue
		}

		if isGroupTaskString(str) {
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

func parseTaskString(str string) *domain.Task {
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

	task := domain.NewTask(title, isDone)
	return task
}

func parseGroupTaskTitle(str string) *domain.Group {
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := domain.NewGroup(str)
	return g
}

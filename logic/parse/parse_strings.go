package parse

import (
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

func StringsToTasks(taskStrings []string) *domain.TodoList {
	todoItemContainer := domain.NewTodoList()
	taskStrings, _ = formatTaskStrings(taskStrings)
	var group *domain.Group
	for _, str := range taskStrings {
		if isSingleTaskString(str) {
			task := toTask(str)
			todoItemContainer.AddTask(task)
			continue
		}

		if isGroupTitle(str) {
			group = toGroup(str)
			todoItemContainer.AddGroup(group)
			continue
		}

		if isGroupTaskString(str) {
			str = strings.TrimSpace(str)
			task := toTask(str)
			if group != nil {
				group.AddTask(task)
			}
			continue
		}
	}
	return todoItemContainer
}

func toTask(str string) *domain.Task {
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

func toGroup(str string) *domain.Group {
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := domain.NewGroup(str)
	return g
}
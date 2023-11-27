package parse

import (
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
	"github.com/MayugeStudio/easy_task/logic/format"
	"github.com/MayugeStudio/easy_task/logic/internal/share"
)

func ToTodoList(taskStrings []string) *domain.TodoList {
	list := domain.NewTodoList()
	taskStrings, _ = format.ToValidStrings(taskStrings)
	var group *domain.Group
	currentIndentLevel := 0
	for _, str := range taskStrings {
		if share.IsSingleTaskString(str) {
			currentIndentLevel = 0
			task := toTask(str)
			list.AddItem(task)
			continue
		}

		if share.IsGroupTitle(str) {
			indentLevel := share.GetIndentLevel(str)
			if group != nil && indentLevel > currentIndentLevel {
				currentIndentLevel = indentLevel
				nextGroup := toGroup(str)
				group.AddItem(nextGroup)
				group = nextGroup
				continue
			}
			group = toGroup(str)
			list.AddItem(group)
			continue
		}

		if share.IsGroupTaskString(str) {
			str = strings.TrimSpace(str)
			task := toTask(str)
			if group != nil {
				group.AddItem(task)
			}
			continue
		}
	}
	return list
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
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "-")
	str = strings.TrimSpace(str)
	g := domain.NewGroup(str)
	return g
}

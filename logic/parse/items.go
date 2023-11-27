package parse

import (
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
	"github.com/MayugeStudio/easy_task/logic/internal/share"
)

func ToItems(taskStrings []string) domain.Items {
	items := domain.NewItems()
	var group *domain.Group
	currentIndentLevel := 0
	for _, str := range taskStrings {
		if share.IsSingleTaskString(str) {
			currentIndentLevel = 0
			task := toTask(str)
			items = append(items, task)
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
			items = append(items, group)
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
	return items
}

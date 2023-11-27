package parse

import (
	"github.com/MayugeStudio/easy_task/domain"
	"github.com/MayugeStudio/easy_task/utils"
)

func toTask(str string) *domain.Task {
	title := ""
	isDone := false
	l := utils.NewLine(str).
		TrimSpace().
		TrimPrefix("-").
		TrimSpace().
		Replace("[", "", 1).
		Replace("]", "", 1).
		TrimSpace()

	if l.HasPrefix("X") {
		isDone = true
		l = l.TrimPrefix("X").TrimSpace()
	}

	title = l.String()

	task := domain.NewTask(title, isDone)
	return task
}

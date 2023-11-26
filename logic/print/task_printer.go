package print

import (
	"cmp"
	"fmt"
	"io"
	"slices"

	"github.com/MayugeStudio/easy_task/domain"
)

func Tasks(w io.Writer, tasks []*domain.Task) error {
	maxLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		taskString := getTaskString(task, maxLength)
		if _, err := fmt.Fprintln(w, taskString); err != nil {
			return fmt.Errorf("printing task: %w", err)
		}
	}
	return nil
}

func getTaskString(task *domain.Task, maxLength int) string {
	var doneStr string
	if task.IsDone {
		doneStr = doneSymbol
	} else {
		doneStr = undoneSymbol
	}
	return fmt.Sprintf("[%s] %-*s", doneStr, maxLength, task.Title)
}

func getMaxTaskNameLength(tasks []*domain.Task) int {
	longestTitleTask := slices.MaxFunc(tasks, func(a, b *domain.Task) int {
		return cmp.Compare(len(a.Title), len(b.Title))
	})
	return len(longestTitleTask.Title)
}

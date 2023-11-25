package logic

import (
	"fmt"
	"io"
	"strings"

	"github.com/MayugeStudio/easy_task/domain"
)

const (
	doneSymbol               = "X"
	undoneSymbol             = " "
	progressSymbol           = "#"
	defaultProgressBarLength = 40.0
)

func PrintTasks(w io.Writer, tasks []*domain.Task) error {
	maxLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		taskString := getTaskString(task, maxLength)
		if _, err := fmt.Fprintln(w, taskString); err != nil {
			return fmt.Errorf("printing task: %w", err)
		}
	}
	return nil
}

func PrintGroups(w io.Writer, groups []*domain.Group) error {
	for _, group := range groups {
		groupString := getGroupString(group)
		if _, err := fmt.Fprint(w, groupString); err != nil {
			return err
		}
	}
	return nil
}

func PrintProgress(w io.Writer, todoList *domain.TodoList) error {
	progress := todoList.Progress()
	progressString := getProgressString(progress, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, progressString); err != nil {
		return err
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

func getGroupString(group *domain.Group) string {
	var b strings.Builder
	maxLength := getMaxTaskNameLength(group.Tasks)
	titleString := fmt.Sprintf("%s\n", group.Title)
	b.WriteString(titleString)

	for _, task := range group.Tasks {
		taskString := fmt.Sprintf("  %s\n", getTaskString(task, maxLength))
		b.WriteString(taskString)
	}
	progress := group.Progress()
	taskProgressString := fmt.Sprintf("  %s\n", getProgressString(progress, 20))
	b.WriteString(taskProgressString)
	return b.String()
}

func getProgressString(progress, length float64) string {
	barLength := int(progress * length)
	doneStr := strings.Repeat(progressSymbol, barLength)
	undoneStr := strings.Repeat(" ", int(length)-barLength)
	return fmt.Sprintf("[%s%s]%.1f%%", doneStr, undoneStr, progress*100)
}

func getMaxTaskNameLength(tasks []*domain.Task) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

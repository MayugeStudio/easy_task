package logic

import (
	"easy_task/src/domain"
	"fmt"
	"io"
	"strings"
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
		if _, err := fmt.Fprintf(w, "%s\n", taskString); err != nil {
			return err
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

func PrintProgress(w io.Writer, tasks []*domain.Task) error {
	taskProgressString := getTaskProgressString(tasks, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, taskProgressString); err != nil {
		return err
	}
	return nil
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
	taskProgressString := fmt.Sprintf("  %s\n", getTaskProgressString(group.Tasks, 20))
	b.WriteString(taskProgressString)
	return b.String()
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
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

func getTaskProgressString(tasks []*domain.Task, length float64) string {
	taskNum := float64(len(tasks))
	if taskNum == 0 {
		progressBar := strings.Repeat(" ", int(length))
		return fmt.Sprintf("[%s]%d%%", progressBar, 0)
	}
	doneTaskNum := 0.0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	doneTaskRatio := doneTaskNum / taskNum
	doneTaskStrLength := int(doneTaskRatio * length)
	doneTaskStr := strings.Repeat(progressSymbol, doneTaskStrLength)
	undoneTaskStr := strings.Repeat(" ", int(length)-doneTaskStrLength)
	return fmt.Sprintf("[%s%s]%d%%", doneTaskStr, undoneTaskStr, int(doneTaskRatio*100))
}

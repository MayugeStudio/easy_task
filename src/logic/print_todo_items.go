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

func PrintProgress(w io.Writer, todoItems *domain.TodoItemContainer) error {
	progress := calculateProgress(todoItems)
	progressString := getProgressString(progress, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, progressString); err != nil {
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
	progress := calculateGroupProgress(group)
	taskProgressString := fmt.Sprintf("  %s\n", getProgressString(progress, 20))
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

func getProgressString(progress float64, length float64) string {
	if progress == 0 {
		progressBar := strings.Repeat(" ", int(length))
		return fmt.Sprintf("[%s]%d%%", progressBar, 0)
	}
	barLength := int(progress * length)
	doneStr := strings.Repeat(progressSymbol, barLength)
	undoneStr := strings.Repeat(" ", int(length)-barLength)
	return fmt.Sprintf("[%s%s]%d%%", doneStr, undoneStr, int(progress*100))
}

func calculateProgress(todoItems *domain.TodoItemContainer) float64 {
	taskProgress := calculateTaskProgress(todoItems.GetTasks())
	sumOfGroupProgress := 0.0
	groups := todoItems.GetGroups()
	if len(groups) == 0 {
		return taskProgress
	}

	for _, group := range todoItems.GetGroups() {
		progress := calculateGroupProgress(group)
		sumOfGroupProgress += progress
	}

	// Calculate the average of the group progress
	groupProgressAverage := sumOfGroupProgress / float64(len(todoItems.GetGroups()))
	return (taskProgress + groupProgressAverage) / 2
}

func calculateTaskProgress(tasks []*domain.Task) float64 {
	taskNum := float64(len(tasks))
	if taskNum == 0 {
		return 0
	}
	doneTaskNum := 0.0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	return doneTaskNum / taskNum
}

func calculateGroupProgress(group *domain.Group) float64 {
	return calculateTaskProgress(group.Tasks)
}

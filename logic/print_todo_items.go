package logic

import (
	domain2 "easy_task/domain"
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

func PrintTasks(w io.Writer, tasks []*domain2.Task) error {
	maxLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		taskString := getTaskString(task, maxLength)
		if _, err := fmt.Fprintln(w, taskString); err != nil {
			return fmt.Errorf("printing task: %w", err)
		}
	}
	return nil
}

func PrintGroups(w io.Writer, groups []*domain2.Group) error {
	for _, group := range groups {
		groupString := getGroupString(group)
		if _, err := fmt.Fprint(w, groupString); err != nil {
			return err
		}
	}
	return nil
}

func PrintProgress(w io.Writer, todoItems *domain2.TodoList) error {
	progress := calculateProgress(todoItems)
	progressString := getProgressString(progress, defaultProgressBarLength)
	if _, err := fmt.Fprint(w, progressString); err != nil {
		return err
	}
	return nil
}

func getTaskString(task *domain2.Task, maxLength int) string {
	var doneStr string
	if task.IsDone {
		doneStr = doneSymbol
	} else {
		doneStr = undoneSymbol
	}
	return fmt.Sprintf("[%s] %-*s", doneStr, maxLength, task.Title)
}

func getGroupString(group *domain2.Group) string {
	var b strings.Builder
	maxLength := getMaxTaskNameLength(group.Tasks)
	titleString := fmt.Sprintf("%s\n", group.Title)
	b.WriteString(titleString)

	for _, task := range group.Tasks {
		taskString := fmt.Sprintf("  %s\n", getTaskString(task, maxLength))
		b.WriteString(taskString)
	}
	progress := calculateTaskProgress(group.Tasks)
	taskProgressString := fmt.Sprintf("  %s\n", getProgressString(progress, 20))
	b.WriteString(taskProgressString)
	return b.String()
}

func getProgressString(progress, length float64) string {
	barLength := int(progress * length)
	doneStr := strings.Repeat(progressSymbol, barLength)
	undoneStr := strings.Repeat(" ", int(length)-barLength)
	return fmt.Sprintf("[%s%s]%d%%", doneStr, undoneStr, int(progress*100))
}

func getMaxTaskNameLength(tasks []*domain2.Task) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

func calculateProgress(items *domain2.TodoList) float64 {
	tasks := items.GetTasks()
	groups := items.GetGroups()

	if len(tasks) == 0 && len(groups) == 0 {
		return 0
	}

	tasks = append(tasks, flattenGroupTasks(groups)...)
	doneTaskNum := calculateDoneTaskNum(tasks)
	sumOfTasks := len(tasks)

	return float64(doneTaskNum) / float64(sumOfTasks)
}

func calculateTaskProgress(tasks []*domain2.Task) float64 {
	if len(tasks) == 0 {
		return 0
	}
	doneTaskNum := calculateDoneTaskNum(tasks)
	return float64(doneTaskNum) / float64(len(tasks))
}

func calculateDoneTaskNum(tasks []*domain2.Task) int {
	if len(tasks) == 0 {
		return 0
	}
	doneTaskNum := 0
	for _, task := range tasks {
		if task.IsDone {
			doneTaskNum++
		}
	}
	return doneTaskNum
}

func flattenGroupTasks(groups []*domain2.Group) []*domain2.Task {
	result := make([]*domain2.Task, 0)
	for _, group := range groups {
		tasks := group.Tasks
		result = append(result, tasks...)
	}
	return result
}

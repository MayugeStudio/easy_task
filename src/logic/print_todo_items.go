package logic

import (
	"easy_task/src/domain"
	"fmt"
	"io"
)

const (
	DoneSymbol   = "X"
	UndoneSymbol = " "
)

func PrintTasks(w io.Writer, tasks []*domain.Task) error {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		taskString := getTaskString(task, maxTaskNameLength)
		if _, err := fmt.Fprint(w, taskString); err != nil {
			return err
		}
	}
	return nil
}


func getTaskString(task *domain.Task, maxLength int) string {
	var doneStr string
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	return fmt.Sprintf("[%s] %-*s\n", doneStr, maxLength, task.Title)
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

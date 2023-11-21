package src

import (
	"fmt"
	"io"
)

const (
	DoneSymbol   = "X"
	UndoneSymbol = " "
)

func PrintTasks(w io.Writer, tasks []*Task) error {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		if err := printTask(w, task, maxTaskNameLength); err != nil {
			return err
		}
	}
	return nil
}

func printTask(w io.Writer, task *Task, maxTaskNameLength int) error {
	var doneStr string
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	if _, err := fmt.Fprintf(w, "[%s] %-*s\n", doneStr, maxTaskNameLength, task.Title); err != nil {
		return err
	}
	return nil
}

func getMaxTaskNameLength(tasks []*Task) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

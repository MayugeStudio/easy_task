package code

import (
	"fmt"
	"io"
)

func PrintTasks(w io.Writer, tasks []TaskPtr) error {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		if err := printTask(w, task, maxTaskNameLength); err != nil {
			return err
		}
	}
	return nil
}

func printTask(w io.Writer, task TaskPtr, maxTaskNameLength int) error {
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

func getMaxTaskNameLength(tasks []TaskPtr) int {
	maxLength := 0
	for _, task := range tasks {
		if len(task.Title) > maxLength {
			maxLength = len(task.Title)
		}
	}
	return maxLength
}

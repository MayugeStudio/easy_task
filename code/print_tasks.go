package code

import "fmt"

func PrintTasks(tasks []TaskPtr) {
	maxTaskNameLength := getMaxTaskNameLength(tasks)
	for _, task := range tasks {
		printTask(task, maxTaskNameLength)
	}
}

func printTask(task TaskPtr, maxTaskNameLength int) {
	var doneStr string
	if task.IsDone {
		doneStr = DoneSymbol
	} else {
		doneStr = UndoneSymbol
	}
	fmt.Printf("[%s] %-*s\n", doneStr, maxTaskNameLength, task.Title)
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

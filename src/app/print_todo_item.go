package app

import (
	"easy_task/src/logic"
	"fmt"
	"io"
)

func PrintTodoItem(w io.Writer, fileName string) (string, int) {
	lines, scanErr := logic.ScanFile(fileName, logic.FileScanner{})
	if scanErr != nil {
		return fmt.Sprintf("scanning file: %v\n", scanErr), 1
	}
	todoItemContainer := logic.ParseStringsToTasks(lines)
	if err := logic.PrintTasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Sprintf("printing tasks: %v\n", err), 1
	}
	if err := logic.PrintTaskProgress(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Sprintf("printing task progress: %v\n", err), 1
	}
	return "", 0
}

package app

import (
	"easy_task/src/logic"
	"fmt"
	"io"
)

func PrintTodoItem(w io.Writer, fileName string, reader logic.FileReader) error {
	lines, scanErr := logic.ScanFile(fileName, reader)
	if scanErr != nil {
		return fmt.Errorf("scanning file: %w", scanErr)
	}
	todoItemContainer := logic.ParseStringsToTasks(lines)
	if err := logic.PrintTasks(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing tasks: %w", err)
	}
	if err := logic.PrintTaskProgress(w, todoItemContainer.GetTasks()); err != nil {
		return fmt.Errorf("printing task progress: %w", err)
	}
	return nil
}
